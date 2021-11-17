package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.URL.RawQuery)
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("OK"))
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {

	mu := http.NewServeMux()

	mu.HandleFunc("/get", GetHandler)

	mu.HandleFunc("/form", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("FORM表单")
		_ = request.ParseForm()
		name := request.FormValue("name")
		password := request.FormValue("password")
		if name == "zhangsan" && password == "123456" {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("OK"))
		} else {
			http.Error(writer, "账户密码不正确", http.StatusInternalServerError)
		}
	})

	mu.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("JSON")
		bytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		var user User
		err = json.Unmarshal(bytes, &user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if user.Name == "zhangsan" && user.Password == "123456" {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("OK"))
		} else {
			http.Error(writer, "账户密码不正确", http.StatusInternalServerError)
		}
	})

	fmt.Println("服务器运行于 8080 端口")
	log.Fatal(http.ListenAndServe(":8080", mu))
}
