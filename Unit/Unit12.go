package main

import (
	"fmt"
	"time"
)

//使用指针修改值

// 交换函数
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}

func main() {
	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y)

	fmt.Println(int32(time.Now().Unix()))

	//	代码说明如下：
	//
	//    第 6 行，定义一个交换函数，参数为 a、b，类型都为 *int 指针类型。
	//    第 9 行，取指针 a 的值，并把值赋给变量 t，t 此时是 int 类型。
	//    第 12 行，取 b 的指针值，赋给指针 a 指向的变量。注意，此时*a的意思不是取 a 指针的值，而是“a 指向的变量”。
	//    第 15 行，将 t 的值赋给指针 b 指向的变量。
	//    第 21 行，准备 x、y 两个变量，分别赋值为 1 和 2，类型为 int。
	//    第 24 行，取出 x 和 y 的地址作为参数传给 swap() 函数进行调用。
	//    第 27 行，交换完毕时，输出 x 和 y 的值。
}
