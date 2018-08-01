package scheduler

import "cz1y.com/learngo/zhibo8/model"

type SimpleScheduler struct {
	WorkChan chan model.Request
}

func (s *SimpleScheduler) Submit(r model.Request) {
	go func() { s.WorkChan <- r }()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan model.Request) {
	s.WorkChan = c
}
