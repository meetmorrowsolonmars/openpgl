package settings

import (
	"google.golang.org/grpc"

	desc "github.com/meetmorrowsolonmars/openpgl/settings/internal/pb/api/v1/settings"
)

type Implementation struct {
	desc.UnimplementedSettingsServiceServer
}

func NewSettingsServiceImplementation() *Implementation {
	return &Implementation{}
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv desc.SettingsServiceServer) {
	desc.RegisterSettingsServiceServer(s, srv)
}
