package main

import "fmt"

//Go语言指针
//指针地址、指针类型和指针取值
func main() {

	//每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
	//Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作），格式如下：
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p", &cat, &str)

	//	代码说明如下：
	//
	//    第 8 行，声明整型变量 cat。
	//    第 9 行，声明字符串变量 str。
	//    第 10 行，使用 fmt.Printf 的动词%p打印 cat 和 str 变量的内存地址，指针的值是带有0x十六进制前缀的一组数

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	//	代码说明如下：
	//
	//    第 10 行，准备一个字符串并赋值。
	//    第 13 行，对字符串取地址，将指针保存到变量 ptr 中。
	//    第 16 行，打印变量 ptr 的类型，其类型为 *string。
	//    第 19 行，打印 ptr 的指针地址，地址每次运行都会发生变化。
	//    第 22 行，对 ptr 指针变量进行取值操作，变量 value 的类型为 string。
	//    第 25 行，打印取值后 value 的类型。
	//    第 28 行，打印 value 的值。

}
