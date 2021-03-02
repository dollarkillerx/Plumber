package mq_manager

import (
	"fmt"

	"github.com/dollarkillerx/plumber/internal/config"
	"github.com/dollarkillerx/plumber/internal/kafka"
	"github.com/pkg/errors"
)

func init() {
	MQManager.RegisterMQ(config.Kafka, &kafka.Kafka{})
}

var MQManager = &mqManager{
	mqs: map[string]MQ{},
}

type mqManager struct {
	mqs map[string]MQ
	mq  MQ
}

func (m *mqManager) RegisterMQ(mn config.MQEngine, mq MQ) {
	m.mqs[string(mn)] = mq
}

func (m *mqManager) GetMQ() MQ {
	return m.mq
}

func (m *mqManager) InitMQManager(cfg config.BaseConfig) error {
	var mq MQ
	var ex bool

	switch cfg.MQEngine {
	case config.Kafka:
		if cfg.KafkaConfig == nil {
			return errors.New("No configuration for kafka was found")
		}

		mq, ex = m.mqs[config.Kafka.String()]
		if !ex {
			return errors.New("not found kafka plugin")
		}
	case config.NSQ:
		if cfg.NSQConfig == nil {
			return errors.New("No configuration for NSQ was found")
		}

		mq, ex = m.mqs[config.Kafka.String()]
		if !ex {
			return errors.New("not found nsq plugin")
		}
	case config.RabbitMQ:
		if cfg.NSQConfig == nil {
			return errors.New("No configuration for RabbitMQ was found")
		}

		mq, ex = m.mqs[config.Kafka.String()]
		if !ex {
			return errors.New("not found RabbitMQ plugin")
		}
	default:
		return errors.WithStack(fmt.Errorf("not found %s", cfg.MQEngine))
	}

	m.mq = mq
	return mq.InitMQ(cfg)
}
