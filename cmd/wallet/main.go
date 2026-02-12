package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/primal-host/wallet/internal/config"
	"github.com/primal-host/wallet/internal/endpoint"
	"github.com/primal-host/wallet/internal/server"
)

func main() {
	slog.Info("wallet starting", "version", config.Version)

	cfg := config.Load()

	store, err := endpoint.NewStore(cfg.EndpointsFile)
	if err != nil {
		slog.Error("endpoints load failed", "error", err)
		os.Exit(1)
	}
	slog.Info("endpoints loaded", "count", len(store.List()))

	srv := server.New(store, cfg.ListenAddr)

	go func() {
		if err := srv.Start(); err != nil {
			slog.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	slog.Info("shutting down", "signal", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("shutdown error", "error", err)
	}
	slog.Info("stopped")
}
