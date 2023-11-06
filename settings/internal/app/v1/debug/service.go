package debug

import (
	"google.golang.org/grpc"

	desc "github.com/meetmorrowsolonmars/openpgl/settings/internal/pb/api/v1/debug"
	"github.com/meetmorrowsolonmars/openpgl/settings/internal/pkg/services"
)

type Implementation struct {
	desc.UnimplementedDebugServiceServer

	settingsService *services.SettingsService
}

func NewDebugServiceImplementation(settingsService *services.SettingsService) *Implementation {
	return &Implementation{
		settingsService: settingsService,
	}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.DebugServiceServer) {
	desc.RegisterDebugServiceServer(s, srv)
}
