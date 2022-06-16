package main

import (
	"fmt"
	"time"
)

//Go语言计算函数执行时间
func main() {
	test()
	test2()
}

//【示例】使用 Since() 函数获取函数的运行时间。
func test() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("test()该函数执行完成耗时：", elapsed)
}

//【示例 2】使用 time.Now().Sub() 获取函数的运行时间。
func test2() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Now().Sub(start)
	fmt.Println("test2()该函数执行完成耗时：", elapsed)
}
