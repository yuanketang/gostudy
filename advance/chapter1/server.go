package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	fmt.Println("路由和路由处理器")

// 	// 第一种
// 	// 注册路由 "/" 注册路由的path, 默认路径 http://www.baidu.com
// 	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
// 		// 处理应用逻辑

// 		fmt.Println(r.Host)         // 主机名
// 		fmt.Println(r.Header)       // 请求头
// 		fmt.Println(r.Method)       // 请求方法 GET POST
// 		fmt.Println(r.URL.Path)     // Path
// 		fmt.Println(r.URL.RawQuery) // query-string
// 	})

// 	// http.ResponseWriter 负责响应数据的
// 	// http.Request 处理请求数据的
// 	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
// 		rw.WriteHeader(http.StatusOK)
// 		rw.Write([]byte("Hello world!"))
// 	})

// 	http.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
// 		r.ParseForm() // 非常重要

// 		// GET请求
// 		if r.Method == http.MethodGet {
// 			fmt.Println(r.Form.Get("name"))
// 			fmt.Println(r.Form.Get("password"))
// 		}

// 		// POST请求
// 		if r.Method == http.MethodPost {
// 			fmt.Println(r.FormValue("name"))
// 			fmt.Println(r.FormValue("password"))
// 		}
// 	})

// 	// 启动web服务
// 	// 表示绑定所有可有本地接口的8080端口
// 	addr := ":8080"
// 	http.ListenAndServe(addr, nil)
// }

type MyHandler struct {
	Id int
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Printf("你好，世界, %d", h.Id)))
}

func main() {
	// // 第一种方式
	// http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.WriteHeader(http.StatusOK)
	// 	rw.Write([]byte("Hello world!"))
	// })

	// // 第二种方式
	// http.Handle("/hello", &MyHandler{})

	mu := http.NewServeMux()
	mu.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello world!"))
	})

	// 启动服务

	http.ListenAndServe(":8080", mu)
}
