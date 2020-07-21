package book

import "fmt"

type BookDetail struct {
	Name        string
	Author      string
	Publisher   string
	PublishTime string
	Price       string
	Score       string
	Intro       string
}

func (book BookDetail) String() string {
	return fmt.Sprintf("\n名字：《%v》\n作者：%v\n出版社: %v\n出版时间：%v\n价格：%v\n评分：%v\n简介：%v\n", book.Name, book.Author, book.Publisher, book.PublishTime, book.Price, book.Score, book.Intro)
}
