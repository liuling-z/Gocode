package main

import "fmt"

//创建指针的另一种方法——new() 函数
func main() {
	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)
	fmt.Println(str)
}
