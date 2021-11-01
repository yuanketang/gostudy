package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type User struct {
	Id   int    `json:"id" xml:"id,attr"`
	Name string `json:"name" xml:"name,attr"`
}

type JsonParser string

func (p JsonParser) Serialize(v interface{}) {
	if bytes, err := json.Marshal(v); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", bytes)
	}
}

type XmlParser string

func (p XmlParser) Serialize(v interface{}) {
	if bytes, err := xml.Marshal(v); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", bytes)
	}
}

func main() {
	fmt.Println("接口")

	user := User{
		Id:   101,
		Name: "小李",
	}

	var p1 JsonParser
	p1.Serialize(user)

	var p2 XmlParser
	p2.Serialize(user)

	fmt.Println("接口之后")
	printAny(user, p1)
	fmt.Println()
	printAny(user, p2)
}

type TokenParser interface {
	Serialize(v interface{})
}

func printAny(v interface{}, p TokenParser) {
	p.Serialize(v)
}
