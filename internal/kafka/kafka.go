package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/dollarkillerx/plumber/pkg/models"
	"github.com/dollarkillerx/plumber/pkg/newsletter"
	"github.com/pkg/errors"
)

type Kafka struct {
	producer     sarama.SyncProducer
	eventChannel chan *models.MQEvent
	config       newsletter.TaskConfig
}

func (k *Kafka) InitMQ(config newsletter.TaskConfig) error {
	kafkaConfig := sarama.NewConfig()

	if config.KafkaConfig.EnableSASL {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext
		kafkaConfig.Net.SASL.User = config.KafkaConfig.User
		kafkaConfig.Net.SASL.Password = config.KafkaConfig.Password
	}

	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer(config.KafkaConfig.Brokers, kafkaConfig)
	if err != nil {
		return errors.WithStack(err)
	}
	k.producer = producer
	k.config = config
	go k.core()
	return nil
}

func (k *Kafka) core() {
	k.eventChannel = make(chan *models.MQEvent, 1000)
loop:
	for {
		select {
		case mg, ex := <-k.eventChannel:
			if !ex {
				if err := k.producer.Close(); err != nil {
					log.Println(err)
				}
				break loop
			}

			if mg.Table == nil {
				continue
			}

			marshal, err := json.Marshal(mg)
			if err != nil {
				log.Println(err)
				continue
			}

			if _, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
				Topic: k.config.KafkaConfig.Topic,
				Key:   sarama.ByteEncoder(fmt.Sprintf("%s_%s", mg.Table.DBName, mg.Table.TableName)),
				Value: sarama.ByteEncoder(marshal),
			}); err != nil {
				log.Println(err)
			}
		}
	}
}

func (k *Kafka) SendMQ(event *models.MQEvent) error {
	k.eventChannel <- event
	return nil
}

func (k *Kafka) Close() {
	close(k.eventChannel)
}
