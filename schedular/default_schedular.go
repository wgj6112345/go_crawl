package schedular

import "github.com/wgj6112345/go_crawl/model"

type DefaultSchedular struct {
	workChan chan model.Request
	Fetcher  FetchFunc
}

func (s *DefaultSchedular) Run() {
	s.workChan = make(chan model.Request, 10000)
}

func (s *DefaultSchedular) Dispatch(req model.Request) {
	go func() {
		s.workChan <- req
	}()
}

func (s *DefaultSchedular) GetWorkChan() chan model.Request {
	return s.workChan
}

func (s *DefaultSchedular) WorkerIdle(w chan model.Request) {
	return
}
