package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/timut2/music-library-api/config"
	httphandl "github.com/timut2/music-library-api/internal/delivery/http"
	"github.com/timut2/music-library-api/pkg/logger"
)

type httpServer struct {
	handler *httphandl.Handler
	config  *config.Config
}

func NewServer(handler *httphandl.Handler, cfg *config.Config) httpServer {
	return httpServer{
		handler: handler,
		config:  cfg,
	}
}

func (s httpServer) Start() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Port),
		Handler: s.handler.Routes(),
	}

	logger.PrintInfo("starting server", map[string]any{
		"addr": srv.Addr,
	})
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("can't start server")
	}

	return nil
}
