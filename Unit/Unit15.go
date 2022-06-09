package main

import (
	"fmt"
	"math"
	"time"
)

//语言常量、iota 常量生成器
func main() {
	const Pi = 3.14159 //相当于math.Pi的近似值
	//批量声明常量
	const (
		e  = 2.7182818
		pi = 3.1415926
	)

	const (
		a = 1
		b //省略初始化表达式，自动取上初始化表达式
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay) // "time.Duration
	// 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	//iota 常量生成器
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

}
