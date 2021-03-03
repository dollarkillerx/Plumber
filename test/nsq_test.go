package test

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
)

func TestNSQ(t *testing.T) {
	producer, err := nsq.NewProducer("192.168.88.11:4150", nsq.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}

	if err := producer.Publish("T1", []byte("Sssss")); err != nil {
		log.Fatalln(err)
	}
}

func TestNSQConsumer(t *testing.T) {
	consumer, err := nsq.NewConsumer("test1", "tc", nsq.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}

	consumer.AddHandler(&cs{})

	if err := consumer.ConnectToNSQD("192.168.88.11:4150"); err != nil {
		log.Fatalln(err)
	}

	for {
		select {}
	}
}

type cs struct {
}

func (c *cs) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}
