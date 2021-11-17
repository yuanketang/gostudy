package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {

	request, err := http.NewRequest("GET", "http://localhost:8080/get", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 创建一个Recorder记录响应

	recorder := httptest.NewRecorder()

	GetHandler(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("期望返回状态码为 %d 实际返回状态码 %d", http.StatusOK, recorder.Code)
	}

	expected := "OK"
	if recorder.Body.String() != expected {
		t.Errorf("期望返回结果为 %s 实际返回结果为 %s", expected, recorder.Body.String())
	}
}
