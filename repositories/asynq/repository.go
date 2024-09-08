package asynq

import (
	"encoding/json"
	"errors"
	"github.com/rakib-09/golang-clean-architecture/config"
	"github.com/rakib-09/golang-clean-architecture/domain"
	"github.com/rakib-09/golang-clean-architecture/utils/logger"
	"time"

	"github.com/hibiken/asynq"
)

type Repository struct {
	config    *config.AsynqConfig
	client    *asynq.Client
	inspector *asynq.Inspector
}

func NewRepository(config *config.AsynqConfig, client *asynq.Client, inspector *asynq.Inspector) *Repository {
	return &Repository{
		config:    config,
		client:    client,
		inspector: inspector,
	}
}

func (repo *Repository) CreateTask(event domain.AsyncTaskType, data interface{}) (*asynq.Task, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(event.String(), payload), nil
}

func (repo *Repository) EnqueueTask(task *asynq.Task, customOpts *domain.AsynqOption) (string, error) {
	opts := repo.asynqOptions(*customOpts)
	taskInfo, err := repo.client.Enqueue(task, opts...)
	if err != nil {
		return "", err
	}
	return taskInfo.ID, nil
}

func (repo *Repository) DequeueTask(taskID string) error {
	err := repo.inspector.DeleteTask(repo.config.Queue, taskID)
	if errors.Is(err, asynq.ErrTaskNotFound) || errors.Is(err, asynq.ErrQueueNotFound) {
		return nil
	}
	if err != nil {
		logger.Error(err, "error on deleting task: ", taskID)
		return err
	}

	return nil
}

func (repo *Repository) asynqOptions(customOpts domain.AsynqOption) []asynq.Option {
	retryOpt := asynq.MaxRetry(0)
	queueOpt := asynq.Queue(repo.config.Queue)
	retentionOpt := asynq.Retention(repo.config.Retention * time.Hour)

	if customOpts.Retry > 0 {
		retryOpt = asynq.MaxRetry(customOpts.Retry)
	}

	if customOpts.Queue != "" {
		queueOpt = asynq.Queue(customOpts.Queue)
	}

	if customOpts.RetentionSeconds > 0 {
		retentionOpt = asynq.Retention(customOpts.RetentionSeconds * time.Hour)
	}

	opts := []asynq.Option{
		retryOpt,
		queueOpt,
		retentionOpt,
	}

	// zero value not allowed
	if len(customOpts.TaskID) > 0 {
		opts = append(opts, asynq.TaskID(customOpts.TaskID))
	}

	// zero value not allowed
	if customOpts.DelaySeconds > 0 {
		opts = append(opts, asynq.ProcessIn(customOpts.DelaySeconds*time.Second))
	}

	// zero value not allowed
	if customOpts.UniqueTTLSeconds > 0 {
		opts = append(opts, asynq.Unique(customOpts.UniqueTTLSeconds*time.Second))
	}

	return opts
}
