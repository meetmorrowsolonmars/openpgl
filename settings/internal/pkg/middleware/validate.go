package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Validator interface {
	Validate(msg proto.Message) error
}

func ValidateUnaryServerInterceptor(validator Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		msg, ok := req.(proto.Message)
		if !ok {
			return handler(ctx, req)
		}

		err := validator.Validate(msg)
		if err != nil {
			// TODO: add option to handle custom errors
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return handler(ctx, req)
	}
}
