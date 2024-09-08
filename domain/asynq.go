package domain

import (
	"time"

	"github.com/hibiken/asynq"
)

type (
	AsynqOption struct {
		TaskID           string
		Retry            int
		Queue            string
		RetentionSeconds time.Duration
		DelaySeconds     time.Duration
		UniqueTTLSeconds time.Duration
	}

	AsynqRepository interface {
		CreateTask(event AsyncTaskType, paylaod interface{}) (*asynq.Task, error)
		EnqueueTask(task *asynq.Task, customOpts *AsynqOption) (string, error)
		DequeueTask(taskID string) error
	}

	AsynqUseCase interface {
		//CreateStoreGoLiveTask(store *Store, payload *StoreGoLiveTaskPayload) error
	}

	AsyncTaskType string
)

func (at AsyncTaskType) String() string {
	return string(at)
}
