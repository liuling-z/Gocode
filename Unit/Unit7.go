package main

import "fmt"

//数值比较
func main() {
	//布尔值可以和 &&（AND）和 ||（OR）操作符结合，并且有短路行为，
	//如果运算符左边的值已经可以确定整个布尔表达式的值，
	//那么运算符右边的值将不再被求值，因此下面的表达式总是安全的：
	var s string
	fmt.Println(s != "" && s[0] == 'x')

}
