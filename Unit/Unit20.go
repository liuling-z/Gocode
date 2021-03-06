package main

import "fmt"

//Go语言数组详解
func main() {
	//Go语言数组的声明
	var a [3]int             // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	//    语法说明如下所示：
	//
	//    数组变量名：数组声明及使用时的变量名。
	//    元素数量：数组的元素数量，可以是一个表达式，但最终通过编译期计算的结果必须是整型数值，元素数量不能含有到运行时才能确认大小的数值。
	//    Type：可以是任意基本类型，包括数组本身，类型为数组本身时，可以实现多维数组。

	//默认情况下，数组的每个元素都会被初始化为元素类型对应的零值，
	//对于数字类型来说就是 0，同时也可以使用数组字面值语法，用一组值来初始化数组：
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	//在数组的定义中，如果在数组长度的位置出现“...”省略号，
	//则表示数组的长度是根据初始化值的个数来计算，因此，上面数组 q 的定义可以简化为：
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	//数组的长度是数组类型的一个组成部分，因此 [3]int 和 [4]int 是两种不同的数组类型，
	//数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。
	x := [4]int{1, 2, 3}
	x = [4]int{1, 2, 3, 4} // 编译错误：无法将 [4]int 赋给 [3]int
	fmt.Println(x)

	//比较两个数组是否相等
	t := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(t == b, t == c, b == c) // "true false false"
	d := [2]int{1, 2}                   //d := [3]int{1, 2}
	fmt.Println(t == d)                 // 编译错误：无法比较 [2]int == [3]int
}
