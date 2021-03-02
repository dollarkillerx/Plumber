package mq_manager

import (
	"github.com/dollarkillerx/plumber/internal/config"
	"github.com/dollarkillerx/plumber/pkg/models"
)

type MQ interface {
	InitMQ(config.BaseConfig) error
	SendMQ(event *models.MQEvent) error
}
