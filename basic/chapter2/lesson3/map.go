package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Map")

	// map是一种key,value的数据结构
	// map, slice, func
	// map2 := map[string]int{ "a": 5, "b": 4 }
	map1 := map[string]int{
		"a": 5,
		"b": 4,
		"c": 3,
		"d": 2,
		"e": 1,
	}
	// map取值
	fmt.Println(map1["a"])

	// 数组取值
	arr := [3]int{1, 2, 3}
	fmt.Println(arr[0])

	// 数组 遍历
	// for index, val := range arr {
	// 	fmt.Println(index, val)
	// }

	// Map 遍历
	// 知识点 map遍历无序的
	fmt.Println("Map的正序遍历")
	for key, val := range map1 {
		fmt.Println(key, val)
	}

	fmt.Println("Map的正序遍历")
	s := []string{}
	for key := range map1 {
		s = append(s, key)
	}

	sort.Strings(s)
	for _, v := range s {
		fmt.Println(map1[v])
	}

	// make
	map3 := make(map[string]float32)
	fmt.Println("map3", map3)

	test := map[string]int{
		"a": 5,
		"b": 4,
		"c": 3,
		"d": 2,
		"e": 1,
	}

	fmt.Println(test["f"])

	val, ok := test["d"]
	fmt.Println(val, ok)

	if val, ok := test["f"]; ok {
		fmt.Println(val, ok)
	} else {
		fmt.Println("键不存在于map")
	}

}
