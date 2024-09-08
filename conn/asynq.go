package conn

import (
	"github.com/rakib-09/golang-clean-architecture/config"

	"github.com/hibiken/asynq"
)

var asyncClient *asynq.Client
var asynqInspector *asynq.Inspector

func InitAsynqClient() {
	asyncClient = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     config.Asynq().RedisAddr,
		DB:       config.Asynq().DB,
		Password: config.Asynq().Pass,
	})
}

func InitAsyncInspector() {
	asynqInspector = asynq.NewInspector(asynq.RedisClientOpt{
		Addr:     config.Asynq().RedisAddr,
		DB:       config.Asynq().DB,
		Password: config.Asynq().Pass,
	})
}

func Asynq() *asynq.Client {
	return asyncClient
}

func AsynqInspector() *asynq.Inspector {
	return asynqInspector
}
