package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/primal-host/wallet/internal/config"
	"github.com/primal-host/wallet/internal/endpoint"
)

func (s *Server) routes() {
	s.echo.GET("/health", s.handleHealth)
	s.echo.GET("/", s.handleDashboard)
	s.echo.GET("/api/status", s.handleStatus)
	s.echo.POST("/api/rpc/:id", s.handleRPC)
}

func (s *Server) handleHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"version": config.Version,
	})
}

func (s *Server) handleDashboard(c echo.Context) error {
	html := strings.ReplaceAll(dashboardHTML, "{{VERSION}}", config.Version)
	return c.HTML(http.StatusOK, html)
}

// handleStatus polls all endpoints and returns their live status.
func (s *Server) handleStatus(c echo.Context) error {
	statuses := s.store.Poll()
	return c.JSON(http.StatusOK, map[string]any{
		"version":   config.Version,
		"endpoints": statuses,
	})
}

// handleRPC proxies a JSON-RPC request to the named endpoint.
func (s *Server) handleRPC(c echo.Context) error {
	id := c.Param("id")

	// Find the endpoint.
	var target *endpoint.Endpoint
	for _, ep := range s.store.List() {
		if ep.ID == id {
			ep := ep
			target = &ep
			break
		}
	}
	if target == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "endpoint not found"})
	}

	// Parse the incoming JSON-RPC request.
	var req struct {
		Method string `json:"method"`
		Params []any  `json:"params"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	result, err := endpoint.RPCCall(target.URL, req.Method, req.Params)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": err.Error()})
	}

	// Return the raw result so the frontend can handle it.
	return c.JSON(http.StatusOK, map[string]json.RawMessage{"result": result})
}
