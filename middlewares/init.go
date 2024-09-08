package middlewares

import (
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Pre(m.RemoveTrailingSlash())
	e.Use(m.LoggerWithConfig(m.LoggerConfig{
		Format:           `${time_custom} ${remote_ip} ${host} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out} "${user_agent}"` + "\n",
		CustomTimeFormat: "2006-01-02T15:04:05.00",
	}))
	e.Use(m.CORS())
	e.Use(m.Secure())
	e.Use(m.Recover())
}
