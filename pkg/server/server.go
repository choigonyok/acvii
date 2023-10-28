package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/choigonyok/acvii/pkg/handler"
	"github.com/choigonyok/acvii/pkg/middleware"
	"github.com/gin-gonic/gin"
)

const (
	clientPort = "8080"
)

type Server struct {
	httpServer *http.Server
	idk        bool
}

var DefaultHandler *http.Handler

// New creates new server
func New(m *middleware.Middleware) *Server {
	srv := newHTTPServer(m)
	return &Server{
		httpServer: srv,
		idk:        true,
	}
}

// newHTTPServer create new http server with given middlewares
func newHTTPServer(m *middleware.Middleware) *http.Server {
	router := gin.Default()

	router.Use(m.Get()...)

	handlers := handler.Get(router)

	srv := &http.Server{
		Addr:         ":" + clientPort,
		Handler:      handlers,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv
}

// Start starts server with gracefully close
func (s *Server) Start() {
	go func() {
		s.httpServer.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	<-ctx.Done()
	log.Println("5 seconds finish")
	log.Println("Server exiting")

}

// Stop stops running server
func (s *Server) Stop() {
	s.httpServer.Close()
}
