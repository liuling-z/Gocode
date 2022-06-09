package main

import "fmt"

//Go语言append()为切片添加元素
func main() {
	var a []int
	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	//不过需要注意的是，在使用 append() 函数为切片动态添加元素时，
	//如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变。

	//切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……，代码如下：
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
	//代码说明如下：
	//
	//    第 1 行，声明一个整型切片。
	//    第 4 行，循环向 numbers 切片中添加 10 个数。
	//    第 5 行，打印输出切片的长度、容量和指针变化，使用函数 len() 查看切片拥有的元素个数，使用函数 cap() 查看切片的容量情况。

	//在切片的开发添加元素
	var x = []int{1, 2, 3}
	x = append([]int{0}, x...)          // 在开头添加1个元素
	x = append([]int{-3, -2, -1}, x...) // 在开头添加1个切片

	//因为 append 函数返回新切片的特性，所以切片也支持链式操作，
	//我们可以将多个 append 操作组合起来，实现在切片中间插入元素：

	//	var y []int
	//	y = append(y[:i], append([]int{x}, y[i:]...)...)       // 在第i个位置插入x
	//	y = append(y[:i], append([]int{1, 2, 3}, y[i:]...)...) // 在第i个位置插入切片
}
