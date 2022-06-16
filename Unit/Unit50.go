package main

import "fmt"

//函数的返回值
func main() {
	a, b := typedTwoValues()
	fmt.Println(typedTwoValues())
	fmt.Println(a, b)

	//调用函数
	//下面是对各个部分的说明：
	//
	//    函数名：需要调用的函数名。
	//    参数列表：参数变量以逗号分隔，尾部无须以分号结尾。
	//    返回值变量列表：多个返回值使用逗号分隔。
	result := addTwoNum(1, 1)
	fmt.Println(result)
}

//1) 同一种类型返回值
func typedTwoValues() (int, int) {
	return 1, 2
}

//2) 带有变量名的返回值
func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

//代码说明如下：
//
//    第 1 行，对两个整型返回值进行命名，分别为 a 和 b。
//    第 3 行和第 4 行，命名返回值的变量与这个函数的布局变量的效果一致，可以对返回值进行赋值和值获取。
//    第 6 行，当函数使用命名返回值时，可以在 return 中不填写返回值列表，如果填写也是可行的，下面代码的执行效果和上面代码的效果一样。
func addTwoNum(x int, y int) int { return x + y }
