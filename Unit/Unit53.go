package main

import "fmt"

//匿名函数用作回调函数
func main() {
	// 使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}

// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

//代码说明如下：
//
//    第 8 行，使用 visit() 函数将整个遍历过程进行封装，当要获取遍历期间的切片值时，只需要给 visit() 传入一个回调参数即可。
//    第 18 行，准备一个整型切片 []int{1,2,3,4} 传入 visit() 函数作为遍历的数据。
//    第 19～20 行，定义了一个匿名函数，作用是将遍历的每个值打印出来。
