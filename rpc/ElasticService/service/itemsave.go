package service

import (
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/model/book"
	"imooc/分布式爬虫项目/demo1/schedular"
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
