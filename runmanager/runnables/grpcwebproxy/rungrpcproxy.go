package grpcwebproxy

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"aduu.dev/utils/runmanager"
	"aduu.dev/utils/web/cors"
)

// GRPCWebService holds all that is necessary to run a grpc web service.
type GRPCWebService struct {
	GrpcServer *grpc.Server
	Address    string
	Debug      bool
}

// Run runs the grpc web service.
func (service *GRPCWebService) Run(m *runmanager.RunManager) {
	wrappedGrpc := grpcweb.WrapServer(service.GrpcServer,
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithAllowNonRootResource(true),
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}))

	// Pass the reques to this router in case it is no grpc web request.
	router := http.NewServeMux()

	combinedHandler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
			return
		}
		router.ServeHTTP(resp, req)
	})

	server := http.Server{
		Addr:         service.Address,
		Handler:      combinedHandler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Set default http additionalHeaders
	server.Handler = defaultHeaders(server.Handler)

	if service.Debug {
		//httpServer.Handler = handlers.LoggingHandler(os.Stdout, httpServer.Handler)
	}

	//log.Fatal(httpServer.ListenAndServeTLS(config.TLSCertPath, config.TLSKeyPath))

	m.Run(&runmanager.Runner{
		Run: func() error {
			fmt.Println("Running grpcweb proxy server on", server.Addr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}

			return nil
		},
		Shutdown: func() error {
			fmt.Println("Shutting down grpcweb proxy server")
			return server.Shutdown(m.Context)
		},
	})
}

type additionalHeaders struct {
	from http.Handler
}

func (h *additionalHeaders) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	h.from.ServeHTTP(resp, req)
}

func defaultHeaders(handler http.Handler) http.Handler {
	return &additionalHeaders{
		from: cors.AllowAllCORSFromEverywhere(handler, "x-grpc-web", "x-user-agent"),
	}
}
