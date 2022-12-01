package appServer

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// Run
// Запуск сервера.
func (s *Server) Run(port string, handler http.Handler) error {

	ADDR := fmt.Sprintf(":%v", port)

	s.httpServer = &http.Server{
		Addr:           ADDR,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	fmt.Printf("\nServer started on %s\n", fmt.Sprintf("http://localhost%s", ADDR))

	err := s.httpServer.ListenAndServe()

	return err
}

// Shutdown
// Остановка сервера.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
