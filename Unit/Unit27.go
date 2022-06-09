package main

import "fmt"

//Go语言range关键字：循环迭代切片
func main() {

	//通过前面的学习我们了解到切片其实就是多个相同类型元素的连续集合，
	//既然切片是一个集合，那么我们就可以迭代其中的元素，
	//Go语言有个特殊的关键字 range，它可以配合关键字 for 来迭代切片里的每一个元素，如下所示：

	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	//【示例 1】range 提供了每个元素的副本
	// 创建一个整型切片，并赋值
	slice = []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}

	//【示例 2】使用空白标识符（下划线）来忽略索引值
	// 创建一个整型切片，并赋值
	slice = []int{10, 20, 30, 40}
	// 迭代每个元素，并显示其值
	for _, value := range slice {
		fmt.Printf("Value: %d\n", value)
	}

	// 【示例 3】使用传统的 for 循环对切片进行迭代
	// 创建一个整型切片，并赋值
	slice = []int{10, 20, 30, 40}
	// 从第三个元素开始迭代每个元素
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
}
