package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fishmanDK/rest-entangle/config"
	"github.com/fishmanDK/rest-entangle/internal/handlers"
	"github.com/fishmanDK/rest-entangle/internal/service"
)

const (
	_json = "json"
	_env = "env"
	_yml = "yml"
)

func main() {
	cfg := config.New(_env)

	_service := service.New(cfg.Url)
	_handlers := handlers.New(_service)

	http.HandleFunc("/getTotalSupply", _handlers.Handler)

	server := http.Server{
		Addr: ":" + cfg.Port,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Server error", slog.String("err", err.Error()))
			os.Exit(1)
		}
	}()

	log.Println("rest-entangle Started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	sig := <-stop
	fmt.Printf("Received signal: %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server shutdown error", slog.String("err", err.Error()))
		os.Exit(1)
	}

	log.Println("Server gracefully stopped")
}