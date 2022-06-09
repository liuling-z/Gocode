package main

import "fmt"

//变量的作用域
//函数外声明为全局变量
var a = 34
var b float32 = 3.14

func main() {
	//函数内声明为局部变量
	//就近原则
	a := 3
	var b int = 4
	c := sum(a, b)
	fmt.Printf("main()中 a = %d, b = %d, c = %d\n\n", a, b, c)

}

//形式参数
func sum(a int, b int) int {
	fmt.Printf("sum()函数中 a=%d\n", a)
	fmt.Printf("sum()函数中 b=%d\n", b)
	num := a + b
	return num
}
