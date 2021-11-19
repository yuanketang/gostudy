package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gostudy/advance/chapter2/lesson1/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	db  *sql.DB
	err error
)

// 打开数据库
func init() {
	// user:password@tcp(host:port)/db[?parameters]
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// 释放数据库连接
	defer db.Close()

	mu := http.NewServeMux()

	mu.HandleFunc("/articles", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			rows, err := db.Query("select id, title, content, created_at from articles order by id DESC limit 10")
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()
			var (
				result    []map[string]interface{}
				id        int
				title     string
				content   string
				createdAt time.Time
			)
			for rows.Next() {
				err = rows.Scan(&id, &title, &content, &createdAt)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusInternalServerError)
					return
				}
				result = append(result, map[string]interface{}{
					"id":         id,
					"title":      title,
					"content":    content,
					"created_at": createdAt.Format("2006-01-02 15:04:05"),
				})
			}
			if rows.Err() != nil {
				http.Error(writer, rows.Err().Error(), http.StatusInternalServerError)
				return
			}
			tpl, err := template.ParseFiles("./advance/chapter1/lesson1/public/articles.html")
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tpl.Execute(writer, result)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		} else {
			http.Error(writer, "只支持GET请求", http.StatusMethodNotAllowed)
		}
	})

	mu.HandleFunc("/article", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			err := utils.OutputHtml(writer, request, "./advance/chapter1/lesson1/public/new.html")
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		} else if request.Method == http.MethodPost {
			_ = request.ParseForm()
			stmt, err := db.Prepare("insert into articles(title, content) values(?, ?)")
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			result, err := stmt.Exec(request.FormValue("title"), request.FormValue("content"))
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			effectRows, err := result.RowsAffected()
			if effectRows == 0 {
				http.Error(writer, "新增文章失败", http.StatusInternalServerError)
				return
			}
			writer.Header().Set("Location", "/articles")
			writer.WriteHeader(http.StatusFound)
		}
	})

	mu.HandleFunc("/delete", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		id, _ := strconv.Atoi(request.FormValue("id"))
		stmt, err := db.Prepare("delete from articles where id = ?")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := stmt.Exec(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		effectRows, err := result.RowsAffected()
		if effectRows == 0 {
			http.Error(writer, "删除文章失败", http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Location", "/articles")
		writer.WriteHeader(http.StatusFound)
	})

	fmt.Println("服务器运行于 8080 端口")
	log.Fatal(http.ListenAndServe(":8080", mu))
}
