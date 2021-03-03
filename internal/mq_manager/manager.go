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
	registerMQ: map[string]MQ{},
}

type mqManager struct {
	registerMQ map[string]MQ // 注册的MQ
	mqs        map[string]MQ // 激活的MQ
}

func (m *mqManager) RegisterMQ(mn config.MQEngine, mq MQ) {
	m.registerMQ[string(mn)] = mq
}

func (m *mqManager) GetMQ(mqs config.MQEngine) (MQ, error) {
	mq, ex := m.mqs[mqs.String()]
	if !ex {
		return nil, errors.WithStack(fmt.Errorf("not found: %s", mq))
	}

	return mq, nil
}

func (m *mqManager) InitMQManager(cfg config.BaseConfig) error {
	var mq MQ
	var ex bool

	switch cfg.MQEngine {
	case config.Kafka:
		if cfg.KafkaConfig == nil {
			return errors.New("No configuration for kafka was found")
		}

		mq, ex = m.registerMQ[config.Kafka.String()]
		if !ex {
			return errors.New("not found kafka plugin")
		}
	case config.NSQ:
		if cfg.NSQConfig == nil {
			return errors.New("No configuration for NSQ was found")
		}

		mq, ex = m.registerMQ[config.Kafka.String()]
		if !ex {
			return errors.New("not found nsq plugin")
		}
	case config.RabbitMQ:
		if cfg.NSQConfig == nil {
			return errors.New("No configuration for RabbitMQ was found")
		}

		mq, ex = m.registerMQ[config.Kafka.String()]
		if !ex {
			return errors.New("not found RabbitMQ plugin")
		}
	default:
		return errors.WithStack(fmt.Errorf("not found %s", cfg.MQEngine))
	}

	m.mq = mq
	return mq.InitMQ(cfg)
}
