package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/choigonyok/goopt/pkg/handler"
	"github.com/choigonyok/goopt/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
	idk        bool
}

var DefaultHandler *http.Handler

func New() *Server {
	srv := newHTTPServer()
	return &Server{
		httpServer: srv,
		idk:        true,
	}
}

func newHTTPServer() *http.Server {
	router := gin.Default()

	middlewares := middleware.New()
	middlewares.AllowOrigin("*")
	middlewares.AllowMethod("GET", "POST", "DELETE", "PUT")
	middlewares.AllowHeader("Origin", "X-Requested-With", "Content-Type", "Accept")
	middlewares.AllowCredential()
	router.Use(middlewares.Get()...)

	handlers := handler.Get(router)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handlers,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv
}

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

func (s *Server) Stop() {
	s.httpServer.Close()
}
