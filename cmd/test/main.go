package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type ss struct {
	Foo int
}

func main() {
	bsonm := bson.M{"foo": 1}
	bytes, _ := bson.Marshal(bsonm)
	fmt.Println(bytes)
	bsonM := ss{}
	err := bson.Unmarshal(bytes, &bsonM)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bsonM.Foo)
}
