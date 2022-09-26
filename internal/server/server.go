package server

import (
	"context"
	"micro_service_phone/pkg/file_logger"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// --------------------------------------------------------------------------------------
func (s *Server) Run(port string, handler http.Handler, logger *file_logger.FileLogger) error {
	logger.Logger.Debug("Runnig server at port:", port)
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// --------------------------------------------------------------------------------------
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
