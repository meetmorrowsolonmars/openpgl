package main

import (
	"context"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	const (
		ServerAddressGRPC  = ":82"
		ServerAddressDebug = ":84"
	)

	ctx := context.Background()
	cancel := func() {}

	// logger setup
	logger := zap.Must(zap.NewProductionConfig().Build()).Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	// metrics setup
	metrics := grpcprom.NewServerMetrics(grpcprom.WithServerHandlingTimeHistogram())
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	// setup db connection
	uri := os.Getenv("MONGODB_URI")
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatalf("can't connect to mongodb: %s", err)
	}

	_ = db

	// grpc server setup
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(metrics.UnaryServerInterceptor()))

	metrics.InitializeMetrics(server)
	reflection.Register(server)

	// debug server startup
	mux := http.NewServeMux()
	mux.Handle("/debug", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	debugServer := http.Server{Addr: ServerAddressDebug, Handler: mux}

	go func() {
		if err := debugServer.ListenAndServe(); err != nil {
			logger.Errorf("debug listen and serve returns error: %s", err)
			cancel()
		}
	}()

	// graceful shutdown
	ctx, cancel = signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		<-ctx.Done()

		// grpc server graceful stop
		server.GracefulStop()

		// debug server graceful stop
		iCtx, iCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer iCancel()
		_ = debugServer.Shutdown(iCtx)
	}()

	// grpc server startup
	logger.Infof("Server startup on port: %s", ServerAddressGRPC)

	lis, err := net.Listen("tcp", ServerAddressGRPC)
	if err != nil {
		logger.Fatalf("can't start app: %s", err)
	}

	if err = server.Serve(lis); err != nil {
		logger.Errorf("grpc server error: %s", err)
		cancel()
	}

	wg.Wait()
}
