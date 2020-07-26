package schedular

import (
	"imooc/分布式爬虫项目/demo1/model"
)

type FetchFunc func(string) ([]byte, error)

type QueueSchedular struct {
	requestChan chan model.Request
	workerChan  chan chan model.Request
	Fetcher     FetchFunc
}

func (s *QueueSchedular) Run() {
	s.requestChan = make(chan model.Request)
	s.workerChan = make(chan chan model.Request)
	go func() {
		var requestQ []model.Request
		var workerQ []chan model.Request

		for {
			var activeReq model.Request
			var activeWorker chan model.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeReq = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case req := <-s.requestChan:
				requestQ = append(requestQ, req)
			case worker := <-s.workerChan:
				workerQ = append(workerQ, worker)
			case activeWorker <- activeReq:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

func (s *QueueSchedular) Dispatch(req model.Request) {
	s.requestChan <- req
}

func (s *QueueSchedular) GetWorkChan() chan model.Request {
	return make(chan model.Request)
}

func (s *QueueSchedular) WorkerIdle(w chan model.Request) {
	s.workerChan <- w
}
