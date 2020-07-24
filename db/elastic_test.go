package db

import (
	"context"
	"fmt"
	"testing"
)

type Hotel struct {
	Location string `json:"location"`
	Price    int    `json:"price"`
}

func TestInsert(t *testing.T) {
	t.Skip()
	var index = "city"
	var typ = "hotel"
	var id = "9928"

	var payload = Hotel{
		Location: "武汉",
		Price:    119,
	}

	fmt.Println("----test insert ----")
	err := Insert(index, typ, id, payload)
	if err != nil {
		fmt.Printf("save item id: %v failed, err : %v\n", id, err)
		return
	}
	fmt.Printf("save item id: %v success\n", id)

	_, err = QueryAll(index, typ)
	if err != nil {
		fmt.Printf("QueryAll item failed, err : %v\n", err)
		return
	}
	fmt.Printf("QueryAll success\n")
}

type Tweet struct {
	User    string
	Message string
}

func TestExample(t *testing.T) {
	// tweet := Tweet{"olivere", "Take Six"}
	// _, err := client.Index().
	// 	Index("twitter").
	// 	Type("tweet").
	// 	Id("23").
	// 	BodyJson(tweet).
	// 	Do(context.Background())
	// if err != nil {
	// 	fmt.Printf("save item failed, err : %v\n", err)
	// 	return
	// }
	exists, err := client.IndexExists("twitter").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("exists： %v\n", exists)
}
