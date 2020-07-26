package book

import "fmt"

type BookDetail struct {
	Cate        string `json:"category"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	PublishTime string `json:"publish_time"`
	Price       string `json:"price"`
	Score       string `json:"score"`
	Intro       string `json:"intro"`
}

func (book BookDetail) String() string {
	return fmt.Sprintf("\n类别：%v \n名字：%v\n作者：%v\n出版社: %v\n出版时间：%v\n价格：%v\n评分：%v\n简介：%v\n", book.Cate, book.Name, book.Author, book.Publisher, book.PublishTime, book.Price, book.Score, book.Intro)
}

type BookItem struct {
	Url        string `json:"url"`
	Type       string `json:"type"`
	Id         string `json:"id"`
	BookDetail `json:"book_detail"`
}

func (book BookItem) String() string {
	return fmt.Sprintf("\nurl: %v \nid: %v \n类别：%v \n名字：《%v》\n作者：%v\n出版社: %v\n出版时间：%v\n价格：%v\n评分：%v\n简介：%v\n", book.Url, book.Id, book.BookDetail.Cate, book.BookDetail.Name, book.BookDetail.Author, book.BookDetail.Publisher, book.BookDetail.PublishTime, book.BookDetail.Price, book.BookDetail.Score, book.BookDetail.Intro)
}
