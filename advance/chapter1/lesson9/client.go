package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {

	//GetData()

	//PostJsonData()

	PostFormData()
}

func GetData() {
	// http://localhost:8080/get?keyword=zhangsan
	params := url.Values{}
	params.Set("keyword", "zhangsan")

	requestUrl, err := url.ParseRequestURI("http://localhost:8080/get")
	if err != nil {
		panic(err)
	}
	requestUrl.RawQuery = params.Encode()
	rawUrl := requestUrl.String()
	fmt.Println(rawUrl)
	resp, err := http.Get(rawUrl)
	if err != nil {
		panic(err)
	}
	dumpResp, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", dumpResp)
}

func PostFormData() {
	params := url.Values{}
	params.Set("name", "zhangsan")
	params.Set("password", "123456")
	resp, err := http.PostForm("http://localhost:8080/form", params)
	if err != nil {
		panic(err)
	}
	dumpResp, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", dumpResp)
}

func PostJsonData() {
	data := `{"name":"zhangsan", "password":"123456"}`
	resp, err := http.Post("http://localhost:8080/json", "application/json", strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	dumpResp, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", dumpResp)
}
