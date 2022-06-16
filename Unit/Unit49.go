package main

import (
	"fmt"
	"math"
)

//Go语言函数（Go语言func）
func main() {
	//Go 语言的函数属于“一等公民”（first-class），也就是说：
	//
	//    函数本身可以作为值进行传递。
	//    支持匿名函数和闭包（closure）。
	//    函数可以满足接口。

	//在Go语言中，函数的基本组成为：关键字 func、函数名、参数列表、返回值、函数体和返回语句，
	//每一个程序都包含很多的函数，函数是基本的代码块。

	//Go语言里面拥三种类型的函数：
	//
	//    普通的带有名字的函数
	//    匿名函数或者 lambda 函数
	//    方法
	fmt.Println(hypot(3, 4))      // "5"
	fmt.Printf("%T\n", add(3, 4)) // "func(int, int) int"
	fmt.Printf("%T\n", sub)       // "func(int, int) int"
	fmt.Printf("%T\n", first)     // "func(int, int) int"
	fmt.Printf("%T\n", zero)      // "func(int, int) int"
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }
