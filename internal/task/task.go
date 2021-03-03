package task

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"time"

	"github.com/dollarkillerx/plumber/internal/mq_manager"
	"github.com/dollarkillerx/plumber/pkg/newsletter"
	"github.com/siddontang/go-mysql/canal"
)

type Task struct {
	mq  mq_manager.MQ
	Cfg newsletter.TaskConfig
	canal.DummyEventHandler
	canal *canal.Canal
}

func New(mq mq_manager.MQ, cfg newsletter.TaskConfig) *Task {
	if cfg.CDCStartTimestamp == 0 {
		cfg.CDCStartTimestamp = time.Now().Unix()
	}

	if mq == nil {
		log.Println("what fuck")
	} else {
		fmt.Println("iiii")
	}

	return &Task{
		mq:  mq,
		Cfg: cfg,
	}
}

// 同步
func (s *Task) Synchronize() error {
	if s.mq == nil {
		log.Fatalln("what fuck", s.Cfg)
	}

	defaultConfig := canal.NewDefaultConfig()
	defaultConfig.Addr = fmt.Sprintf("%s:%d", s.Cfg.DBConfig.Host, s.Cfg.DBConfig.Port)
	defaultConfig.User = s.Cfg.DBConfig.User
	defaultConfig.Password = s.Cfg.DBConfig.Password
	defaultConfig.Flavor = string(s.Cfg.Engine)

	canal, err := canal.NewCanal(defaultConfig)
	if err != nil {
		s.mq.Close()
		return errors.WithStack(err)
	}

	canal.SetEventHandler(s)
	s.canal = canal
	if err := canal.Run(); err != nil {
		log.Println(err)
		s.mq.Close()
		return errors.WithStack(err)
	}

	return nil
}

func (s *Task) Close() error {
	s.canal.Close()
	s.mq.Close()

	return nil
}
