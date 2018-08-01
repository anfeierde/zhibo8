package main

import (
	"cz1y.com/learngo/zhibo8/engine"
	"cz1y.com/learngo/zhibo8/model"
	"cz1y.com/learngo/zhibo8/parser"
	"cz1y.com/learngo/zhibo8/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(model.Request{
		Url:        "https://db.qiumibao.com/f/index/matchs",
		ParserFunc: parser.Word,
	})

}
