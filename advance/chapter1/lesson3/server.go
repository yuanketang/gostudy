package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// 添加标签
type User struct {
	Id       int    `json:"id" validate:"gt=0"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func main() {
	fmt.Println("表单进阶")

	// JSON

	// 初始化路由管理器
	mu := http.NewServeMux()

	// 注册路由
	mu.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {

		// 从请求体中读取数据
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		// 反序列化数据到我们自定义的结构

		var user User
		err = json.Unmarshal(data, &user)
		if err != nil {
			panic(err)
		}

		log.Printf("User = %#v", user)

		validate := validator.New()
		if err = validate.Struct(user); err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("OK"))
		}
	})

	// multipart-data
	mu.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		const maxFileSize = 10 << 20 // 10M
		err := request.ParseMultipartForm(maxFileSize)
		if err != nil {
			panic(err)
		}
		for _, headers := range request.MultipartForm.File {
			for _, file := range headers {
				// 定义上传文件保存的路径
				saveFilePath := filepath.Join("./", file.Filename)

				// 原始上传文件
				src, err := file.Open()
				if err != nil {
					panic(err)
				}
				defer src.Close()

				// 创建一个空文件
				dst, err := os.Create(saveFilePath)
				if err != nil {
					panic(err)
				}
				defer dst.Close()

				// 从原输入拷贝至我们的目标文件
				if _, err = io.Copy(dst, src); err != nil {
					panic(err)
				}
			}
		}

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("上传完成"))
	})

	// 启动服务
	_ = http.ListenAndServe(":8080", mu)
}
