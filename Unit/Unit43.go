package main

import "fmt"

//遍历通道（channel）——接收通道数据
func main() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

	//代码说明如下：
	//
	//    第 1 行创建了一个整型类型的通道。
	//    第 3 行启动了一个 goroutine，其逻辑的实现体现在第 5～8 行，实现功能是往通道中推送数据 1、2、3，然后结束并关闭通道。
	//    这段 goroutine 在声明结束后，在第 9 行马上被执行。
	//    从第 11 行开始，使用 for range 对通道 c 进行遍历，其实就是不断地从通道中取数据，直到通道被关闭。

	//在遍历中选择希望获得的变量
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for _, value := range m {
		fmt.Println(value)
	}

	//在上面的例子中将 key 变成了下划线_，这里的下划线就是匿名变量。
	//
	//    可以理解为一种占位符。
	//    匿名变量本身不会进行空间分配，也不会占用一个变量的名字。
	//    在 for range 可以对 key 使用匿名变量，也可以对 value 使用匿名变量。
	for key, _ := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d \n", key)
	}

	//我们总结一下 for 的功能：
	//
	//    Go语言的 for 包含初始化语句、条件表达式、结束语句，这 3 个部分均可缺省。
	//    for range 支持对数组、切片、字符串、map、通道进行遍历操作。
	//    在需要时，可以使用匿名变量对 for range 的变量进行选取。
}
