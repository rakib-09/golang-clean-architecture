package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rakib-09/golang-clean-architecture/controllers"
	m "github.com/rakib-09/golang-clean-architecture/middlewares"
)

type Routes struct {
	echo *echo.Echo
	myc  *controllers.MyController
}

func New(e *echo.Echo, myc *controllers.MyController) *Routes {
	return &Routes{
		echo: e,
		myc:  myc,
	}
}

func (r *Routes) Init() {
	e := r.echo
	m.Init(e)

	// swagger docs
	// dg := e.Group("docs")
	// dg.GET("/swagger", echo.WrapHandler(m.SwaggerDocs()))
	// dg.GET("/redoc", echo.WrapHandler(m.ReDocDocs()))
	// dg.GET("/rapidoc", echo.WrapHandler(m.RapiDocs()))
	// e.File("/swagger.yaml", "./swagger.yaml")

	g := e.Group("/v1")
	// public

	// internal
	i := g.Group("/internal")

	i.GET("/hello", r.myc.Hello, m.CustomAuth())
}
