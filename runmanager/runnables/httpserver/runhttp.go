package httpserver

import (
	"fmt"
	"net/http"

	"aduu.dev/utils/runmanager"
)

// Runnable impplements a runnable http service.
type Runnable struct {
	Endpoint string
	Handler  http.Handler
}

// Run runs the http server on the specified port to return html.
func (s *Runnable) Run(m *runmanager.RunManager) {
	server := http.Server{
		Addr:    s.Endpoint,
		Handler: s.Handler,
	}

	m.Run(&runmanager.Runner{
		Run: func() error {
			fmt.Println("Running http server on", server.Addr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}

			return nil
		},
		Shutdown: func() error {
			fmt.Println("Shutting down http server")
			return server.Shutdown(m.Context)
		},
	})
}

/*
// Run the http server on the specified port to return html.
func (s *SingleFolderAPI) Run(m *RunManager) {
	server := http.Server{
		Addr: s.HTTPServerAddress(),
		Handler: &httpServerImpl{
			settings: s,
		},
	}

	m.Run(&Runner{
		Run: func() error {
			fmt.Println("Running http server on", server.Addr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}

			return nil
		},
		Shutdown: func() error {
			fmt.Println("Shutting down http server")
			return server.Shutdown(m.Context)
		},
	})
}
*/
