package main

import (
	"fmt"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"net/http"
)

// 需要导入依赖 go get github.com/go-oauth2/oauth2/v4

func main() {

	// 1 manager相关配置
	manager := manage.NewDefaultManager()
	// 设置默认Token过期时间等等参数
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	// 指定Token存储位置
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	// 创建一个Token生成器并纳入manager管理
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	// 配置客户端
	clientStore := store.NewClientStore()
	clientStore.Set("123456", &models.Client{
		ID: "123456",
		Secret: "666666",
		Domain: "http://127.0.0.1:8080",
	})
	// 客户端纳入manager管理
	manager.MapClientStorage(clientStore)

	// 2 启动OAuth认证服务器
	srv := server.NewServer(server.NewConfig(), manager)

	// 是否允许使用GET方式传递Token
	srv.SetAllowGetAccessRequest(true)


	mu := http.NewServeMux()

	// 颁发授权码给客户端
	// http://127.0.0.1:8081/authcode?client_id=123456&redirect_uri=http%3A%2F%2F127.0.0.1%3A8080%2Foauth2&response_type=code
	mu.HandleFunc("/authcode", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()

		// 验证请求
		req, err := srv.ValidationAuthorizeRequest(request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		// 获取Token
		token, err := srv.GetAuthorizeToken(request.Context(), req)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		data := srv.GetAuthorizeData(req.ResponseType, token)

		url, err := srv.GetRedirectURI(req, data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		writer.Header().Set("Location", url)
		writer.WriteHeader(http.StatusFound)
	})

	// 颁发令牌给客户端
	mu.HandleFunc("/token", func(writer http.ResponseWriter, request *http.Request) {
		err := srv.HandleTokenRequest(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
	})

	fmt.Println("服务端端运行在 8081 端口")
	http.ListenAndServe(":8081", mu)
}
