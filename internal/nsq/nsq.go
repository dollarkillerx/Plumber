package nsq

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/dollarkillerx/plumber/pkg/models"
	"github.com/dollarkillerx/plumber/pkg/newsletter"
	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
)

type NSQ struct {
	cfg      newsletter.TaskConfig
	producer *nsq.Producer

	eventChannel chan *models.MQEvent
}

func (n *NSQ) InitMQ(cfg newsletter.TaskConfig) error {
	if len(cfg.NSQConfig.Addr) == 0 {
		return errors.New("nsq addr is null")
	}
	n.cfg = cfg
	if err := n.retry(); err != nil {
		return err
	}

	n.eventChannel = make(chan *models.MQEvent, 1000)
	go n.core()
	return nil
}

func (n *NSQ) retry() error {
	rand.Seed(time.Now().UnixNano())
	addr := n.cfg.NSQConfig.Addr[rand.Intn(len(n.cfg.NSQConfig.Addr))]
	producer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		return errors.WithStack(err)
	}

	n.producer = producer

	return nil
}

func (n *NSQ) SendMQ(event *models.MQEvent) error {
	n.eventChannel <- event
	return nil
}

func (n *NSQ) Close() {
	close(n.eventChannel)
}

func (n *NSQ) core() {
loop:
	for {
		select {
		case mg, ex := <-n.eventChannel:
			if !ex {
				n.producer.Stop()
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
				if err := n.producer.Publish(n.cfg.NSQConfig.Topic, marshal); err != nil {
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
