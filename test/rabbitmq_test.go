package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/streadway/amqp"
)

func TestRabbitMQ(t *testing.T) {
	dial, err := amqp.Dial("amqp://admin:admin@127.0.0.1:5672/")
	if err != nil {
		log.Fatalln(err)
	}

	defer dial.Close()
	channel, err := dial.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"test1",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatalln(err)
	}

	consume, err := channel.Consume(q.Name, "s1", true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

loop:
	for {
		select {
		case e, ex := <-consume:
			if !ex {
				break loop
			}

			fmt.Println(string(e.Body))
		}
	}
}
