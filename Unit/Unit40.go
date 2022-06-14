package main

import "fmt"

//for 中的初始语句——开始循环时执行的语句
func main() {
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}
	//for 中的条件表达式——控制是否循环的开关
	//1) 结束循环时带可执行语句的无限循环
	var i int
	for ; ; i++ {
		if i > 10 {
			break
		}
	}
	//2) 无限循环
	for {
		if i > 10 {
			break
		}
		i++
	}

	//3) 只有一个循环条件的循环
	for i <= 10 {
		i++
	}

	// for 中的结束语句——每次循环结束时执行的语句
	//在结束每次循环前执行的语句，如果循环被 break、goto、return、panic 等语句强制退出，结束语句不会被执行。
}
