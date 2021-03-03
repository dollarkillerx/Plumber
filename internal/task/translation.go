package task

import (
	"github.com/dollarkillerx/plumber/internal/utils"
	"github.com/siddontang/go-mysql/canal"
)

func (s *Task) OnRow(e *canal.RowsEvent) error {
	if e == nil {
		return nil
	}
	if e.Header == nil {
		return nil
	}

	if int64(e.Header.Timestamp) < s.Cfg.CDCStartTimestamp {
		return nil
	}

	event := utils.PkgMQEvent(e)
	return s.mq.SendMQ(event)
}

func (s *Task) String() string {
	return "Plumber"
}
