package service

import (
	"github.com/wgj6112345/go_crawl/logger"
	"github.com/wgj6112345/go_crawl/model/book"
	"github.com/wgj6112345/go_crawl/schedular"
)

type DemoSaveService struct {
}

func (s *DemoSaveService) Save(item book.BookItem, result *string) (err error) {
	logger.Logger.Infof("save url %v\n", item.Url)

	*result, err = schedular.DemoSave(item)
	if err != nil {
		logger.Logger.Errorf("demosave failed, err : %v\n", err)
		*result = "demosave failed"
		return
	}

	return
}
