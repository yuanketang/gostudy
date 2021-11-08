package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"gostudy/advance/chapter1/lesson5/utils"
	"log"
	"net/http"
)

// 配置客户端
var (
	config = oauth2.Config{
		ClientID: "123456",
		ClientSecret: "666666",
		Endpoint: oauth2.Endpoint{
			AuthURL: "http://127.0.0.1:8081/authcode",
			TokenURL: "http://127.0.0.1:8081/token",
		},
		RedirectURL: "http://127.0.0.1:8080/oauth2", // 回调地址
	}
)

func main() {

	mu := http.NewServeMux()

	mu.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := utils.OutputHtml(writer, request, "advance/chapter1/lesson5/static/index.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
		}
	})

	mu.HandleFunc("/auth", func(writer http.ResponseWriter, request *http.Request) {
		err := utils.OutputHtml(writer, request, "advance/chapter1/lesson5/static/auth.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
		}
	})

	// http://127.0.0.1:8081/authcode?client_id=123456&redirect_uri=http%3A%2F%2F127.0.0.1%3A8080%2Foauth2&response_type=code
	mu.HandleFunc("/wechat", func(writer http.ResponseWriter, request *http.Request) {
		url := config.AuthCodeURL("")
		http.Redirect(writer, request, url, http.StatusFound)
	})

	mu.HandleFunc("/oauth2", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		code := request.FormValue("code")
		if code == "" {
			http.Error(writer, "code不正确", http.StatusBadRequest)
			return
		}

		token, err := config.Exchange(request.Context(), code)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		e := json.NewEncoder(writer)
		e.SetIndent("", "  ")
		_ = e.Encode(token)
	})

	mu.HandleFunc("/client", func(writer http.ResponseWriter, request *http.Request) {
		cfg := clientcredentials.Config{
			ClientID: config.ClientID,
			ClientSecret: config.ClientSecret,
			TokenURL: config.Endpoint.TokenURL,
		}
		token, err := cfg.Token(request.Context())
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		e := json.NewEncoder(writer)
		e.SetIndent("", "  ")
		_ = e.Encode(token)
	})

	fmt.Println("客户端运行在 8080 端口")
	log.Fatal(http.ListenAndServe(":8080", mu))
}
