package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/rakib-09/golang-clean-architecture/controllers"
	"github.com/rakib-09/golang-clean-architecture/routes"
	"github.com/rakib-09/golang-clean-architecture/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// Clients
	//redisClient := conn.NewRedisClient()
	//dbClient := conn.Db()
	//asyncClient := conn.Asynq()
	//asyncInspector := conn.AsynqInspector()

	// Repositories
	//dbRepo := db.NewRepository(dbClient)
	//redisRepo := redis.NewRepository(redisClient)
	//asynqRepo := asynq.NewRepository(config.Asynq(), asyncClient, asyncInspector)

	// Services
	//asynqSvc := services.NewAsynqService(config.Asynq(), asynqRepo)

	// Controllers
	myc := controllers.NewController()

	// Server
	var echo_ = echo.New()
	var Routes = routes.New(echo_, myc)
	var Server = server.New(echo_)

	// Spooling
	Routes.Init()
	Server.Start()
}
