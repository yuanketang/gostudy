package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Id   int
	Name string
	Age  int
	Sex  bool
	Addr Address
}

type Address struct {
	Province string
	City     string
	Street   string
}

func main() {

	// 1. 需要解析我们的模板文件
	// 2. 传递数据到我们的模板文件
	// 3. 在模板文件中使用模板语法来显示数据

	mu := http.NewServeMux()

	// 简单使用
	mu.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 1. 需要解析我们的模板文件
		tpl, err := template.ParseFiles("./advance/chapter1/lesson7/public/index.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		// 2. 传递数据到我们的模板文件
		err = tpl.Execute(writer, false)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	})

	// 结构体
	mu.HandleFunc("/struct", func(writer http.ResponseWriter, request *http.Request) {
		// 1. 需要解析我们的模板文件
		tpl, err := template.ParseFiles("./advance/chapter1/lesson7/public/struct.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		// 2. 传递数据到我们的模板文件
		user := User{
			Id:   666,
			Name: "小李",
			Age:  20,
			Addr: Address{
				Province: "广东",
				City:     "广州",
				Street:   "123456",
			},
		}
		err = tpl.Execute(writer, user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	})

	// Map
	mu.HandleFunc("/map", func(writer http.ResponseWriter, request *http.Request) {
		// 1. 需要解析我们的模板文件
		tpl, err := template.ParseFiles("./advance/chapter1/lesson7/public/map.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		// 2. 传递数据到我们的模板文件
		err = tpl.Execute(writer, map[string]interface{}{
			"id":    1,
			"price": 50.0,
			"qty":   2,
			"total": 100.0,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	})

	// Map2
	mu.HandleFunc("/map2", func(writer http.ResponseWriter, request *http.Request) {
		// 1. 需要解析我们的模板文件
		tpl, err := template.ParseFiles("./advance/chapter1/lesson7/public/map2.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		// 2. 传递数据到我们的模板文件
		err = tpl.Execute(writer, []map[string]interface{}{
			{
				"id":    1,
				"price": 50.0,
				"qty":   2,
				"total": 100.0,
			},
			{
				"id":    2,
				"price": 99.0,
				"qty":   1,
				"total": 99.0,
			},
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	})

	// 条件
	// with false情况：数值0 ，false，指针nil, 数组，切片len() == 0
	mu.HandleFunc("/if", func(writer http.ResponseWriter, request *http.Request) {
		// 1. 需要解析我们的模板文件
		tpl, err := template.ParseFiles("./advance/chapter1/lesson7/public/if.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		// 2. 传递数据到我们的模板文件
		user := User{
			Id:   666,
			Name: "小李",
			Age:  20,
			Sex:  false,
			Addr: Address{
				Province: "广东",
				City:     "广州",
				Street:   "123456",
			},
		}
		err = tpl.Execute(writer, user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	})

	fmt.Println("服务器运行于 8080 端口")
	log.Fatal(http.ListenAndServe(":8080", mu))
}
