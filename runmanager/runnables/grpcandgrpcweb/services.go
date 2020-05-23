package grpcandgrpcweb

import (
	"google.golang.org/grpc"

	"aduu.dev/utils/runmanager"
	"aduu.dev/utils/runmanager/runnables/grpcserver"
	"aduu.dev/utils/runmanager/runnables/grpcwebproxy"
)

// RegisterGrpcServiceFunc registers a grpc endpoint with the given grpc server.
type RegisterGrpcServiceFunc = func(server *grpc.Server)

// Services creates grpc and grpcweb service functions to use with a runmanager setup.
func Services(grpcEndpoint string, grpcWebEndpoint string,
	registerFunc RegisterGrpcServiceFunc) []runmanager.Service {
	grpcService := grpcserver.Runnable{
		Endpoint:            grpcEndpoint,
		RegisterGRPCService: registerFunc,
	}

	return []runmanager.Service{
		func(m *runmanager.RunManager) {
			grpcService.Run(m)
		}, func(m *runmanager.RunManager) {
			webService := &grpcwebproxy.GRPCWebService{
				GrpcServer: grpcService.GRPCServer,
				Address:    grpcWebEndpoint,
			}

			webService.Run(m)
		},
	}
}
