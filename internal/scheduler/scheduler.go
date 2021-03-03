package scheduler

import (
	"github.com/dollarkillerx/plumber/internal/config"
	"github.com/dollarkillerx/plumber/internal/task"
)

type Scheduler struct {
	tasks map[string]*task.Task
	cfg   config.BaseConfig
}

func New(cfg config.BaseConfig) *Scheduler {
	return &Scheduler{
		cfg:   cfg,
		tasks: map[string]*task.Task{},
	}
}
