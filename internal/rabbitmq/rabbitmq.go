package rabbitmq

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/dollarkillerx/plumber/pkg/models"
	"github.com/dollarkillerx/plumber/pkg/newsletter"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	cfg     newsletter.TaskConfig
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   *amqp.Queue

	eventChannel chan *models.MQEvent
}

func (n *RabbitMQ) InitMQ(cfg newsletter.TaskConfig) error {
	n.cfg = cfg
	if err := n.retry(); err != nil {
		return err
	}

	n.eventChannel = make(chan *models.MQEvent, 1000)
	go n.core()
	return nil
}

func (n *RabbitMQ) retry() error {
	rand.Seed(time.Now().UnixNano())
	dial, err := amqp.Dial(n.cfg.RabbitMQConfig.Uri)
	if err != nil {
		return errors.WithStack(err)
	}
	channel, err := dial.Channel()
	if err != nil {
		return errors.WithStack(err)
	}

	declare, err := channel.QueueDeclare(
		n.cfg.RabbitMQConfig.Queue,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return errors.WithStack(err)
	}

	n.conn = dial
	n.channel = channel
	n.queue = &declare
	return nil
}

func (n *RabbitMQ) SendMQ(event *models.MQEvent) error {
	n.eventChannel <- event
	return nil
}

func (n *RabbitMQ) Close() {
	close(n.eventChannel)
}

func (n *RabbitMQ) core() {
loop:
	for {
		select {
		case mg, ex := <-n.eventChannel:
			if !ex {
				if err := n.channel.Close(); err != nil {
					log.Println(err)
				}
				if err := n.conn.Close(); err != nil {
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

			for i := 0; i < 3; i++ {
				if err := n.channel.Publish("", n.queue.Name, false, false,
					amqp.Publishing{DeliveryMode: amqp.Persistent, ContentType: "application/json", Body: marshal}); err != nil {
					if err := n.retry(); err != nil {
						log.Printf("%+v\n", err)
					}
					continue
				}
				break
			}
		}
	}
}
