package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run (port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown (ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
