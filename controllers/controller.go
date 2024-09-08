package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type MyController struct {
}

func NewController() *MyController {
	return &MyController{}
}

func (mc MyController) Hello(c echo.Context) error {

	return c.JSON(http.StatusOK, "hello world")
}
