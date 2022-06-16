package main

import "fmt"

//Go语言可变参数（变参函数）
func main() {
	myfunc(2, 3, 4)
	myfunc(1, 3, 7, 13)
	myfunc2([]int{2, 3, 4})
	myfunc2([]int{1, 3, 7, 13})

}

//可变参数类型
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func myfunc2(args []int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
