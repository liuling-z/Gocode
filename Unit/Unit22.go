package main

import "fmt"

//Go语言多维数组简述
func main() {
	//声明二维数组
	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array)

	//【示例 3】同样类型的多维数组赋值
	// 声明两个二维整型数组
	var array1 [2][2]int
	var array2 [2][2]int
	// 为array2的每个元素赋值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	// 将 array2 的值复制给 array1
	array1 = array2
	fmt.Println(array1)

	//【示例 4】使用索引为多维数组赋值
	// 将 array1 的索引为 1 的维度复制到一个同类型的新数组里
	var _ [2]int = array1[1]
	// 将数组中指定的整型值复制到新的整型变量里
	var _ int = array1[1][0]

}
