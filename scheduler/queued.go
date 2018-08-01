package scheduler

import "cz1y.com/learngo/zhibo8/model"

type QueuedScheduler struct {
	requestChan chan model.Request
	workerChan  chan chan model.Request
}

func (s *QueuedScheduler) Submit(r model.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan model.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) WorkerChan() chan model.Request {
	return make(chan model.Request)
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan model.Request)
	s.requestChan = make(chan model.Request)
	go func() {
		var requestQ []model.Request
		var workerQ []chan model.Request

		for {
			var activityRequest model.Request
			var activityWorker chan model.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activityWorker = workerQ[0]
				activityRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activityWorker <- activityRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
