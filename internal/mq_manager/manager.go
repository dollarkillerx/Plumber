package mq_manager

import (
	"fmt"

	"github.com/dollarkillerx/plumber/internal/kafka"
	"github.com/dollarkillerx/plumber/pkg/newsletter"
	"github.com/pkg/errors"
)

func init() {
	MQManager.RegisterMQ(newsletter.Kafka, &kafka.Kafka{})
}

var MQManager = &mqManager{
	registerMQ: map[string]MQ{},
}

type mqManager struct {
	registerMQ map[string]MQ // 注册的MQ
}

func (m *mqManager) RegisterMQ(mn newsletter.MQEngine, mq MQ) {
	m.registerMQ[string(mn)] = mq
}

func (m *mqManager) InitMQManager(cfg newsletter.TaskConfig) (MQ, error) {
	var mq MQ
	var ex bool

	switch cfg.MQEngine {
	case newsletter.Kafka:
		if cfg.KafkaConfig == nil {
			return nil, errors.New("No configuration for kafka was found")
		}

		mq, ex = m.registerMQ[newsletter.Kafka.String()]
		if !ex {
			return nil, errors.New("not found kafka plugin")
		}
	case newsletter.NSQ:
		if cfg.NSQConfig == nil {
			return nil, errors.New("No configuration for NSQ was found")
		}

		mq, ex = m.registerMQ[newsletter.NSQ.String()]
		if !ex {
			return nil, errors.New("not found nsq plugin")
		}
	case newsletter.RabbitMQ:
		if cfg.NSQConfig == nil {
			return nil, errors.New("No configuration for RabbitMQ was found")
		}

		mq, ex = m.registerMQ[newsletter.RabbitMQ.String()]
		if !ex {
			return nil, errors.New("not found RabbitMQ plugin")
		}
	default:
		return nil, errors.WithStack(fmt.Errorf("not found %s", cfg.MQEngine))
	}

	err := mq.InitMQ(cfg)
	return mq, err
}
