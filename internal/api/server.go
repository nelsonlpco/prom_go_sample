package api

import (
	"github.com/labstack/echo"

	handler2 "github.com/nelsonlpco/classic_cc_problens/internal/api/handler"
)

type Server struct {
	e       *echo.Echo
	handler *handler2.Handler
}

func NewServer() *Server {
	return new(Server)
}

func (s *Server) WithEcho(e *echo.Echo) *Server {
	s.e = e
	return s
}

func (s *Server) WithHandler(h *handler2.Handler) *Server {
	s.handler = h
	return s
}

func (s Server) mapRoutes() {
	s.e.GET(handler2.FibonacciPath, s.handler.FibonacciCalc)
	s.e.POST(handler2.DNAPath, s.handler.CompressDNA)
}

func (s Server) Start() {
	s.mapRoutes()
	s.e.Start(":8080")
}
