package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)

const secret = "123456"

func main() {

	mu := http.NewServeMux()

	mu.HandleFunc("/token", func(writer http.ResponseWriter, request *http.Request) {

		// JWT payload 公共部分
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * 60).Unix(),
			Issuer:    "yuanketang",
		})
		tokenStr, err := token.SignedString([]byte(secret))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		e := json.NewEncoder(writer)
		_ = e.Encode(tokenStr)
	})

	// GET or Post
	mu.HandleFunc("/valid", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			_ = request.ParseForm()
			tokenStr := request.FormValue("token")
			token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				if jwtErr, ok := err.(*jwt.ValidationError); ok {
					if jwtErr.Errors == jwt.ValidationErrorExpired {
						http.Error(writer, "Token过期了", http.StatusBadRequest)
						return
					}
				}
				return
			}
			if !token.Valid {
				http.Error(writer, "认证失败", http.StatusUnauthorized)
				return
			}
			writer.Header().Set("Content-Type", "application/json")
			e := json.NewEncoder(writer)
			e.SetIndent("", "  ")
			_ = e.Encode(token)
		} else {
			// POST
			// Authorization: Bearer xxzcxzczxczxczxcxz
			tokenStr := request.Header.Get("Authorization")
			tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")

			token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				if jwtErr, ok := err.(*jwt.ValidationError); ok {
					if jwtErr.Errors == jwt.ValidationErrorExpired {
						http.Error(writer, "Token过期了", http.StatusBadRequest)
						return
					}
				}
				return
			}
			if !token.Valid {
				http.Error(writer, "认证失败", http.StatusUnauthorized)
				return
			}
			writer.Header().Set("Content-Type", "application/json")
			e := json.NewEncoder(writer)
			e.SetIndent("", "  ")
			_ = e.Encode(token)
		}
	})

	fmt.Println("服务器运行于 8080 端口")
	log.Fatal(http.ListenAndServe(":8080", mu))
}
