package schedular

import "imooc/分布式爬虫项目/demo1/model"

type DefaultSchedular struct {
	WorkChan chan model.Request
}

func (s *DefaultSchedular) Dispatch(req model.Request) {
	go func() {
		s.WorkChan <- req
	}()
}

func (s *DefaultSchedular) Init() {
	s.WorkChan = make(chan model.Request, 10000)
}

func (s *DefaultSchedular) GetWorkChan() chan model.Request {
	return s.WorkChan
}
