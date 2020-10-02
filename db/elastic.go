package db

import (
	"context"
	"fmt"
	"github.com/wgj6112345/go_crawl/logger"
	"github.com/wgj6112345/go_crawl/model/book"
	"reflect"

	"github.com/olivere/elastic/v7"
)

const (
	elastic_addr = "http://127.0.0.1:9200"
)

var (
	client *elastic.Client
)

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(elastic_addr), elastic.SetSniff(false))
	if err != nil {
		logger.Logger.Errorf("connect to elastic failed, err : %v\n", err)
		panic(err)
	}

	info, code, err := client.Ping(elastic_addr).Do(context.Background())
	if err != nil {
		logger.Logger.Errorf("ping elasticsearch failed, err : %v\n", err)
		panic(err)
	}

	logger.Logger.Debugf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	return
}

func Insert(index string, typ string, id string, data interface{}) (err error) {
	fmt.Println("--insert inside--")

	_, err = client.Index().Index(index).Type(typ).Id(id).BodyJson(data).Do(context.Background())
	if err != nil {
		fmt.Println("--insert inside err--")
		logger.Logger.Errorf("insert data  %v to elastic failed, err : %v\n", data, err)
		return
	}
	fmt.Println("--insert inside success--")
	logger.Logger.Infof("insert index %s, type %s, id %s success\n", index, typ, id)
	return
}

func Delete(index string, typ string, id string) (err error) {
	res, err := client.Delete().Index(index).Type(typ).Id(id).Do(context.Background())
	if err != nil {
		logger.Logger.Errorf("delete from elastic failed, err : %v\n", err)
		return
	}
	logger.Logger.Errorf("delete data  %v from elastic success\n", res.Result)
	return
}

func Update(index string, typ string, id string, field string, value interface{}) (err error) {
	_, err = client.Update().Index(index).Type(typ).Id(id).Doc(map[string]interface{}{field: value}).Do(context.Background())
	if err != nil {
		logger.Logger.Errorf("update field %v set value: %v from elastic failed, err : %v\n", field, err)
		return
	}
	logger.Logger.Errorf("update field %v from elastic success\n", field)
	return
}

// get : 通过 id 来查找
func Get(index string, typ string, id string) (get *elastic.GetResult, err error) {
	get, err = client.Get().Index(index).Type(typ).Id(id).Do(context.Background())
	if err != nil {
		logger.Logger.Errorf("get id: %v from elastic failed, err : %v\n", id, err)
		return
	}

	if get.Found {
		logger.Logger.Debugf("Get document %s in version %d from index %s, type %s\n", get.Id, get.Version, get.Index, get.Type)
		return
	}
	return
}

// query: 搜索 满足一定条件的
func QueryAll(index string, typ string) (res *elastic.SearchResult, err error) {
	res, err = client.Search(index).Type(typ).Do(context.Background())
	if err != nil {
		logger.Logger.Errorf("query index: %v type: %v from elastic failed, err : %v\n", index, typ, err)
		return
	}
	printSearchResult(res)
	return
}

func QueryString(index string, typ string, field string, value string) (res *elastic.SearchResult, err error) {
	qString := fmt.Sprintf("%s:%s", field, value)
	q := elastic.NewQueryStringQuery(qString)

	res, err = client.Search(index).Type(typ).Query(q).Do(context.Background())
	if err != nil {
		logger.Logger.Errorf("search %v from elastic failed, err : %v\n", qString, err)
		return
	}
	return
}

// TODO  条件查询  短语搜索    简单分页

func printSearchResult(res *elastic.SearchResult) {
	var bookItem book.BookItem
	for _, item := range res.Each(reflect.TypeOf(bookItem)) {
		t := item.(book.BookItem)
		fmt.Println("t: ", t)
	}
}
