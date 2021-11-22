package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gostudy/advance/chapter2/lesson2/utils"
	"html/template"
	"log"
	"net/http"
	"time"
)


// 连接数据库mongodb

var (
	db *mongo.Client
	collection *mongo.Collection
	err error
)

// 自定义结构体
type Article struct {
	// primitive.ObjectID
	Id		string		`bson:"_id,omitempty"`
	Title 	string		`bson:"title"`
	Content	string		`bson:"content"`
	CreatedAt	string	`bson:"created_at"`
}

func init() {
	db, err = mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://127.0.0.1:27017"),
		)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = db.Database("test").Collection("articles")
}

func main() {

	mu := http.NewServeMux()

	mu.HandleFunc("/articles", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			// 查询数据
			findOpt := options.FindOptions{}
			// 模拟 order by id DESC limit 10
			findOpt.SetLimit(10).SetSort(bson.M{"_id": -1})
			cursor, err := collection.Find(request.Context(), bson.M{}, &findOpt)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			defer cursor.Close(request.Context())

			var result []Article
			err = cursor.All(request.Context(), &result)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			// 处理模板
			tpl, err := template.ParseFiles("./advance/chapter2/lesson2/public/articles.html")
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tpl.Execute(writer, result)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		}
	})

	mu.HandleFunc("/article", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			err := utils.OutputHtml(writer,request,"./advance/chapter2/lesson2/public/new.html")
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		} else if request.Method == http.MethodPost {
			// 解析表单参数
			_ = request.ParseForm()

			article := &Article{
				Title: request.FormValue("title"),
				Content: request.FormValue("content"),
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			}
			_, err = collection.InsertOne(request.Context(), article)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			writer.Header().Set("Location", "/articles")
			writer.WriteHeader(http.StatusFound)
		}
	})

	mu.HandleFunc("/delete", func(writer http.ResponseWriter, request *http.Request) {
		// 解析URL中的参数
		_ = request.ParseForm()

		// 获取ID
		id := request.FormValue("id")
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := collection.DeleteOne(request.Context(), bson.M{"_id": objId})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if result.DeletedCount == 0 {
			http.Error(writer, "删除文章失败", http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Location", "/articles")
		writer.WriteHeader(http.StatusFound)
	})

	fmt.Println("服务器运行于 8080 端口")
	log.Fatal(http.ListenAndServe(":8080", mu))
}
