package server

import (
	"github.com/rakib-09/golang-clean-architecture/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func (s *Server) Start() {
	e := s.echo
	e.Logger.Fatal(e.Start(":" + config.App().Port))
}

func New(echo *echo.Echo) *Server {
	return &Server{echo: echo}
}
