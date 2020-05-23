package grpcproxy

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"aduu.dev/utils/runmanager"
)

// RegisterGRPCService registers a grpc server to proxy through mux.
type RegisterGRPCService func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

// Runnable proxies a grpc service through http.
type Runnable struct {
	Register       RegisterGRPCService
	FromAddress    string
	ToAddress      string
	AllowedOrigins []string
}

// Run runs the given grpc proxy as a runnable service.
func (s *Runnable) Run(m *runmanager.RunManager) {
	// Register gRPC server Endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Registering _myservice_ and setting to which grpc-Endpoint to connect to.
	if err := s.Register(m.Context, mux, s.FromAddress, opts); err != nil {
		m.ErrChan <- err
	}

	server := http.Server{
		Addr: s.ToAddress,
		// Handler: allowAllCORS(mux),
		Handler: s.allowCORSFrom(mux, s.AllowedOrigins),
	}

	m.Run(&runmanager.Runner{
		Run: func() error {
			fmt.Println("Running grpc proxy on", server.Addr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}

			return nil
		},
		Shutdown: func() error {
			fmt.Println("Shutting down grpc proxy")
			return server.Shutdown(m.Context)
		},
	})
}

/*
func (s *SingleFolderAPI) Run(m *RunManager) {
	// Register gRPC server Endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Registering _myservice_ and setting to which grpc-Endpoint to connect to.
	if err := myservice.RegisterYourServiceHandlerFromEndpoint(m.Context, mux, s.GrpcServerEndpoint, opts); err != nil {
		m.ErrChan <- err
	}

	server := http.Server{
		Addr: s.HTTPProxyAddress(),
		// Handler: allowAllCORS(mux),
		Handler: s.allowCORSFRomInternalHTTPServer(mux),
	}

	m.Run(&Runner{
		Run: func() error {
			fmt.Println("Running grpc proxy on", server.Addr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}

			return nil
		},
		Shutdown: func() error {
			fmt.Println("Shutting down grpc proxy")
			return server.Shutdown(m.Context)
		},
	})
}



func (s *SingleFolderAPI) allowCORSFRomInternalHTTPServer(h http.Handler) http.Handler {
	return s.allowCORSFrom(h, s.HTTPServerAddress())
}

func (s *SingleFolderAPI) allowCORSFrom(h http.Handler, allowedOrigin string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				if origin == allowedOrigin {
					preflightHandler(w, r)
					return
				}
			}
		}
		h.ServeHTTP(w, r)
	})
}

*/

func (s *Runnable) allowCORSFrom(h http.Handler, allowedOrigins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				for _, allowedOrigin := range allowedOrigins {
					if origin == allowedOrigin {
						preflightHandler(w, r)
						return
					}
				}
			}
		}
		h.ServeHTTP(w, r)
	})
}

/*
// allowAllCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowAllCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
*/

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))

	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}

	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
}
