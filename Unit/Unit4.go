package main

import "fmt"

//多个变量同时赋值
func main() {
	var a int = 100
	var b int = 200
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)

	//从左向右赋值
	b, a = a, b
	fmt.Println(a, b)

	//_ 匿名变量
	a, _ = b, a
	fmt.Println(a, b)

}

//代码说明如下：
//
//第 1 行，将 IntSlice 声明为 []int 类型。
//第 3 行，为 IntSlice 类型编写一个 Len 方法，提供切片的长度。
//第 4 行，根据提供的 i、j 元素索引，获取元素后进行比较，返回比较结果。
//第 5 行，根据提供的 i、j 元素索引，交换两个元素的值。

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
