package main

import "fmt"

//使用“键值对”初始化结构体
//2) 使用匿名结构体的例子
func main() {
	// 实例化一个匿名结构体
	msg := &struct { // 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)
}

// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n", msg)
}

//代码说明如下：
//
//    第 8 行，定义 printMsgType() 函数，参数为 msg，类型为 *struct{id int data string}，因为类型没有使用 type 定义，所以需要在每次用到的地方进行定义。
//    第 14 行，使用字符串格式化中的%T动词，将 msg 的类型名打印出来。
//    第 20 行，对匿名结构体进行实例化，同时初始化成员。
//    第 21 和 22 行，定义匿名结构体的字段。
//    第 24 和 25 行，给匿名结构体字段赋予初始值。
//    第 28 行，将 msg 传入 printMsgType() 函数中进行函数调用。
