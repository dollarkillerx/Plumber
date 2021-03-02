package server

import (
	"fmt"

	"github.com/dollarkillerx/plumber/internal/config"
	"github.com/dollarkillerx/plumber/internal/mq_manager"
	"github.com/dollarkillerx/plumber/internal/utils"
	"github.com/siddontang/go-mysql/canal"
)

type Server struct {
	mq  mq_manager.MQ
	cfg config.BaseConfig
	canal.DummyEventHandler
}

func New(mq mq_manager.MQ, cfg config.BaseConfig) *Server {
	return &Server{
		mq:  mq,
		cfg: cfg,
	}
}

// 同步
func (s *Server) Synchronize() error {
	defaultConfig := canal.NewDefaultConfig()
	defaultConfig.Addr = fmt.Sprintf("%s:%d", s.cfg.DBConfig.Host, s.cfg.DBConfig.Port)
	defaultConfig.User = s.cfg.DBConfig.User
	defaultConfig.User = s.cfg.DBConfig.Password

	canal, err := canal.NewCanal(defaultConfig)
	if err != nil {
		return err
	}

	canal.SetEventHandler(s)
	return canal.Run()
}

func (s *Server) OnRow(e *canal.RowsEvent) error {
	if e == nil {
		return nil
	}
	if e.Header == nil {
		return nil
	}

	if int64(e.Header.Timestamp) < s.cfg.CDCStartTimestamp {
		return nil
	}

	event := utils.PkgMQEvent(e)
	return s.mq.SendMQ(event)
}

func (s *Server) String() string {
	return "Plumber"
}
