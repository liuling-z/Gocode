package main

import "fmt"

//Go语言type关键字（类型别名）
func main() {
	// 将NewInt定义为int类型
	type NewInt int
	// 将int取一个别名叫IntAlias
	type IntAlias = int
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)
	//代码说明如下：
	//
	//    第 8 行，将 NewInt 定义为 int 类型，这是常见的定义类型的方法，通过 type 关键字的定义，NewInt 会形成一种新的类型，NewInt 本身依然具备 int 类型的特性。
	//    第 11 行，将 IntAlias 设置为 int 的一个别名，使用 IntAlias 与 int 等效。
	//    第 16 行，将 a 声明为 NewInt 类型，此时若打印，则 a 的值为 0。
	//    第 18 行，使用%T格式化参数，打印变量 a 本身的类型。
	//    第 21 行，将 a2 声明为 IntAlias 类型，此时打印 a2 的值为 0。
	//    第 23 行，打印 a2 变量的类型。
}
