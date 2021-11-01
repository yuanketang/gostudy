package test

import (
	"gostudy/chapter8/lesson1/demo"
	"testing"
)

// 性能测试用例名称必须以Benchmark开头,后面呢一般是测试的方法或函数名
func BenchmarkAdd(t *testing.B) {
	// 循环
	for i := 0; i < t.N; i++ {
		demo.Add(88, 11)
	}
}
