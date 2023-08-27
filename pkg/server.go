package server

import (
	"context"
	configuration "fantasy-volleyball-api/appconfig"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type server struct {
	engine *gin.Engine
}

func NewServer(engine *gin.Engine) *server {
	return &server{
		engine: engine,
	}
}

func (s *server) StartHttpServer() {
	server := &http.Server{
		Handler: s.engine,
		Addr:    ":" + configuration.Port,
	}
	go func() {
		gracefulShutdown(server)
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("cannot start server.", err)
		panic("cannot start server")
	}

	fmt.Println("Server is running on port: ", configuration.Port)
}

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown Error: ", err)
	}

	fmt.Println("Server exiting")
}
