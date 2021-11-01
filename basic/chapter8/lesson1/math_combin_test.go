package test

import (
	"gostudy/chapter8/lesson1/demo"
	"testing"
	"time"
)

// 测试文件必须以_test.go结尾

// 测试用例

func testAdd(t *testing.T) {
	t.Parallel()
	<-time.After(time.Second * 3)
	// 准备要测试的数据
	// 表格驱动测试
	data := []struct{ a, b, expect int }{
		{1, 2, 3},
		{2, 2, 4},
		{100, 9, 109},
		{36, 51, 87},
		{55, 22, 77},
	}

	for _, row := range data {
		actual := demo.Add(row.a, row.b)
		if actual != row.expect {
			// t.Fail()
			t.Errorf("输入 a = %d, b = %d, 实际计算结果为 actual = %d, 期望值为 expect = %d", row.a, row.b, actual, row.expect)
		}
	}
}

func testMul(t *testing.T) {
	t.Parallel()
	<-time.After(time.Second * 1)
	// 准备要测试的数据
	// 表格驱动测试
	data := []struct{ a, b, expect int }{
		{3, 3, 9},
		{4, 5, 20},
		{11, 11, 121}, //故意写错
	}

	for _, row := range data {
		actual := demo.Mul(row.a, row.b)
		if actual != row.expect {
			// t.Fail()
			t.Errorf("输入 a = %d, b = %d, 实际计算结果为 actual = %d, 期望值为 expect = %d", row.a, row.b, actual, row.expect)
		}
	}
}

func TestCombin(t *testing.T) {
	// 初始化工作
	t.Run("Add", testAdd)
	t.Run("Mul", testMul)
	// 收尾
}
