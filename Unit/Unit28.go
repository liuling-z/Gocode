package main

import "fmt"

//Go语言多维切片简述
func main() {
	//声明一个二维切片
	var slice [][]int
	//为二维切片赋值
	slice = [][]int{{10}, {100, 200}}
	fmt.Println(slice)

	//【示例】组合切片的切片
	// 声明一个二维整型切片并赋值
	slice = [][]int{{10}, {100, 200}}
	// 为第一个切片追加值为 20 的元素
	slice[0] = append(slice[0], 20)
	slice[0] = append(slice[0], 30)
	slice[1] = append(slice[1], 30)
	fmt.Println(slice)
}
