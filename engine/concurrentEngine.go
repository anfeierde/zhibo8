package engine

import (
	"fmt"
	"log"

	"cz1y.com/learngo/zhibo8/fetch"
	"cz1y.com/learngo/zhibo8/model"
)

type Scheduler interface {
	Submit(model.Request)
	WorkerReady(chan model.Request)
	WorkerChan() chan model.Request
	Run()
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e ConcurrentEngine) Run(seeds ...model.Request) {
	out := make(chan model.ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item %v \n", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan model.ParserResult, s Scheduler) {
	go func() {
		in := make(chan model.Request)
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r model.Request) (model.ParserResult, error) {
	log.Printf("Fetching %s \n", r.Url)
	content, err := fetch.Fetch(r.Url)
	if err != nil {
		return model.ParserResult{}, err
	}
	return r.ParserFunc(content), err
}
