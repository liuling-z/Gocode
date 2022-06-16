package main

import "fmt"

//Go语言匿名函数——没有函数名字的函数
func main() {
	//1) 在定义时调用匿名函数
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	//2) 将匿名函数赋值给变量
	// 将匿名函数体保存到f()中
	f := func(data int) {
		fmt.Println("hello", data)
	}
	// 使用f()调用
	f(100)

}
