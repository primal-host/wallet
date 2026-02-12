package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/primal-host/wallet/internal/endpoint"
)

type Server struct {
	echo  *echo.Echo
	store *endpoint.Store
	addr  string
}

func New(store *endpoint.Store, addr string) *Server {
	s := &Server{
		echo:  echo.New(),
		store: store,
		addr:  addr,
	}
	s.echo.HideBanner = true
	s.echo.HidePort = true
	s.echo.Use(middleware.Recover())
	s.routes()
	return s
}

func (s *Server) Start() error {
	slog.Info("server listening", "addr", s.addr)
	if err := s.echo.Start(s.addr); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
