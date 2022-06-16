package main

import "fmt"

//数字阶乘
func main() {
	var i int = 10
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
