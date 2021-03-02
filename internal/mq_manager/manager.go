package mq_manager

import "github.com/dollarkillerx/plumber/internal/config"

var MQManager = &mqManager{
	mqs: map[string]MQ{},
}

type mqManager struct {
	mqs map[string]MQ
}

func (m *mqManager) InitMQManager(cfg config.BaseConfig) error {
	switch cfg.MQEngine {
	case config.Kafka:

	case config.NSQ:

	case config.RabbitMQ:

	}

	return nil
}
