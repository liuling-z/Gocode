package main

import "fmt"

//Go语言函数变量——把函数作为值保存到变量中
func main() {
	var f func()
	f = fire
	f()

}
func fire() {
	fmt.Println("fire")
}

//代码说明：
//
//    第 7 行，定义了一个 fire() 函数。
//    第 13 行，将变量 f 声明为 func() 类型，此时 f 就被俗称为“回调函数”，此时 f 的值为 nil。
//    第 15 行，将 fire() 函数作为值，赋给函数变量 f，此时 f 的值为 fire() 函数。
//    第 17 行，使用函数变量 f 进行函数调用，实际调用的是 fire() 函数。
