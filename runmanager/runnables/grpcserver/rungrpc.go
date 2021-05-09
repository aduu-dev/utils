package grpcserver

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"aduu.dev/utils/runmanager"
)

// Runnable implements a grpc server as a runnable service.
type Runnable struct {
	Endpoint            string
	RegisterGRPCService func(server *grpc.Server)

	GRPCServer *grpc.Server
}

// Run runs the grpc service on top of the given run manager.
func (s *Runnable) Run(m *runmanager.RunManager) {
	lis, err := net.Listen("tcp", s.Endpoint)
	if err != nil {
		m.ErrChan <- err
		return
	}

	address := lis.Addr()

	s.GRPCServer = grpc.NewServer()
	s.RegisterGRPCService(s.GRPCServer)

	m.Run(&runmanager.Runner{
		Run: func() error {
			fmt.Println("Running grpc on", address)

			// Workaround: Print it later again so parent process can pick it up.
			time.AfterFunc(time.Second*2, func() {
				fmt.Println("Running grpc on", address)
			})
			return s.GRPCServer.Serve(lis)
		},
		Shutdown: func() error {
			fmt.Println("Shutting down grpc")
			s.GRPCServer.GracefulStop()
			return nil
		},
	})
}

/*
func (s *SingleFolderAPI) Run(m *RunManager) {
	lis, err := net.Listen("tcp", s.GrpcServerEndpoint)
	if err != nil {
		m.ErrChan <- err
		return
	}

	GRPCServer := grpc.NewServer()
	myservice.RegisterYourServiceServer(GRPCServer, &grpcServerImpl{})

	m.Run(&Runner{
		Run: func() error {
			fmt.Println("Running grpc on", s.GrpcServerEndpoint)
			return GRPCServer.Serve(lis)
		},
		Shutdown: func() error {
			fmt.Println("Shutting down grpc")
			GRPCServer.GracefulStop()
			return nil
		},
	})
}
*/
