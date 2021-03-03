package mq_manager

import (
	"github.com/dollarkillerx/plumber/pkg/models"
	"github.com/dollarkillerx/plumber/pkg/newsletter"
)

type MQ interface {
	InitMQ(cfg newsletter.TaskConfig) error
	SendMQ(event *models.MQEvent) error
	Close()
}
