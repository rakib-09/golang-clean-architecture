package worker

import (
	"fmt"
	"github.com/rakib-09/golang-clean-architecture/config"
	"time"

	"github.com/hibiken/asynq"
)

func Start(mux *asynq.ServeMux) {
	worker := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     config.Asynq().RedisAddr,
			DB:       config.Asynq().DB,
			Password: config.Asynq().Pass,
		},
		asynq.Config{
			Concurrency: config.Asynq().Concurrency,
			Queues: map[string]int{
				config.Asynq().Queue: 1,
			},
			RetryDelayFunc: func(numOfRetry int, e error, t *asynq.Task) time.Duration {
				switch t.Type() {
				default:
					return asynq.DefaultRetryDelayFunc(numOfRetry, e, t)
				}
			},
		},
	)

	if err := worker.Run(mux); err != nil {
		panic(fmt.Sprintf("could not run worker: %v", err))
	}
}
