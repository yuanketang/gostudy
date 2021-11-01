package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	age    int    `json:"age"`
}

func main() {
	fmt.Println("结构体")

	user := new(User)
	user.Id = 999
	user.Name = "小李"
	user.Avatar = "https://www.baidu.com"
	user.age = 20

	bytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", bytes)

	// 反序列化
	fmt.Println("反序列化")
	jsonStr := `{"id":888,"name":"小孙","avatar":"https://www.bilibili.com"}`

	user2 := User{}

	err = json.Unmarshal([]byte(jsonStr), &user2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n, %d, %s, %s", user2, user2.Id, user2.Name, user2.Avatar)
}
