package asynq

import (
	"github.com/rakib-09/golang-clean-architecture/config"
	"github.com/rakib-09/golang-clean-architecture/domain"
)

type AsynqService struct {
	config    config.AsynqConfig
	asynqRepo domain.AsynqRepository
}

func NewAsynqService(config *config.AsynqConfig, asynqRepo domain.AsynqRepository) *AsynqService {
	return &AsynqService{
		config:    *config,
		asynqRepo: asynqRepo,
	}
}
