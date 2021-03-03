package server

import (
	"github.com/dollarkillerx/plumber/internal/utils"
	"github.com/siddontang/go-mysql/canal"
)

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
