package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	mu := http.NewServeMux()

	// compare
	mu.HandleFunc("/compare", func(writer http.ResponseWriter, request *http.Request) {
		// 1. 需要解析我们的模板文件
		tpl, err := template.ParseFiles("./advance/chapter1/lesson8/public/compare.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
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
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	// func
	mu.HandleFunc("/func", func(writer http.ResponseWriter, request *http.Request) {
		// 1. 需要解析我们的模板文件
		tpl, err := template.New("func.html").Funcs(template.FuncMap{
			"discount": func(total float64, rate float64) float64 {
				return total * rate
			},
		}).ParseFiles("./advance/chapter1/lesson8/public/func.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// 2. 传递数据到我们的模板文件
		err = tpl.Execute(writer, map[string]interface{}{
			"id":    1,
			"price": 50.0,
			"qty":   2,
			"total": 240.0,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	// template
	mu.HandleFunc("/template", func(writer http.ResponseWriter, request *http.Request) {
		// 1. 需要解析我们的模板文件
		tpl, err := template.ParseFiles(
			"./advance/chapter1/lesson8/public/template.html",
			"./advance/chapter1/lesson8/public/header.html",
			"./advance/chapter1/lesson8/public/list.html",
		)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// 2. 传递数据到我们的模板文件
		err = tpl.Execute(writer, map[string]interface{}{
			"id":    1,
			"total": 240.0,
			"goods": []map[string]interface{}{
				{"id": 1, "name": "IPhone4", "qty": 1, "price": 4000.0},
				{"id": 2, "name": "IPhone5", "qty": 2, "price": 4200.0},
				{"id": 3, "name": "IPhone6", "qty": 1, "price": 5000.0},
			},
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("服务器运行于 8080 端口")
	log.Fatal(http.ListenAndServe(":8080", mu))
}
