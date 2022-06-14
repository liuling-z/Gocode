package main

import "fmt"

//Go语言for range（键值循环）
func main() {
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	//遍历字符串——获得字符
	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	//遍历 map——获得 map 的键和值
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}
