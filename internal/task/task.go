package task

import (
	"fmt"

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
	return &Task{
		mq:  mq,
		Cfg: cfg,
	}
}

// 同步
func (s *Task) Synchronize() error {
	defaultConfig := canal.NewDefaultConfig()
	defaultConfig.Addr = fmt.Sprintf("%s:%d", s.Cfg.DBConfig.Host, s.Cfg.DBConfig.Port)
	defaultConfig.User = s.Cfg.DBConfig.User
	defaultConfig.User = s.Cfg.DBConfig.Password
	defaultConfig.Flavor = string(s.Cfg.Engine)

	canal, err := canal.NewCanal(defaultConfig)
	if err != nil {
		s.mq.Close()
		return err
	}

	canal.SetEventHandler(s)
	s.canal = canal
	if err := canal.Run(); err != nil {
		s.mq.Close()
		return err
	}

	return nil
}

func (s *Task) Close() error {
	s.canal.Close()
	s.mq.Close()

	return nil
}
