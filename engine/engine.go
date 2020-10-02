package engine

import "github.com/wgj6112345/go_crawl/model"

type Engine interface {
	Run(seeds ...model.Request)
}
