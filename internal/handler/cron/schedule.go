package cron

import (
	"github.com/robfig/cron/v3"
)

type Task interface {
	Handle()
}

type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		cron: cron.New(),
	}
}

func (s *Scheduler) Register(spec string, task Task) error {
	_, err := s.cron.AddFunc(spec, task.Handle)

	return err
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}

func RegisterSchedule(fetchPriceTask *FetchPriceTask) (*Scheduler, error) {
	scheduler := NewScheduler()

	err := scheduler.Register("* * * * *", fetchPriceTask)
	if err != nil {
		return nil, err
	}

	scheduler.Start()

	return scheduler, nil
}
