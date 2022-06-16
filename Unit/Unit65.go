package main

import "fmt"

//Go语言递归函数
func main() {
	//构成递归需要具备以下条件：
	//
	//    一个问题可以被拆分成多个子问题；
	//    拆分前的原问题与拆分后的子问题除了数据规模不同，但处理问题的思路是一样的；
	//    不能无限制的调用本身，子问题需要有退出递归状态的条件。

	//注意：编写递归函数时，一定要有终止条件，否则就会无限调用下去，直到内存溢出。

	//示例：斐波那契数列
	//数列的形式如下所示：
	//
	//1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …
	result := 0
	for i := 1; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}
func fibonacci(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
