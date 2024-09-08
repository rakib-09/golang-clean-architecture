package cmd

import (
	"github.com/rakib-09/golang-clean-architecture/worker"

	asynq_ "github.com/hibiken/asynq"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use: "worker",
	Run: runWorker,
}

func runWorker(cmd *cobra.Command, args []string) {
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
	//Controller := controllers.NewController()

	// Spooling
	mux := asynq_.NewServeMux()
	//mux.HandleFunc(domain.AsyncTaskTypeGoLiveStore.String(), asynqController.ProcessStoreGoLiveTask)

	worker.Start(mux)
}
