package main

import "fmt"

//Go语言switch case语句
func main() {
	//Go语言的 switch 要比C语言的更加通用，表达式不需要为常量，甚至不需要为整数，case 按照从上到下的顺序进行求值，直到找到匹配的项，
	//如果 switch 没有表达式，则对 true 进行匹配，因此，可以将 if else-if else 改写成一个 switch。

	//基本写法
	var a = "hello"
	a = ""
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	a = "mum"
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	}

	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	//跨越 case 的 fallthrough——兼容C语言的 case 设计
	//在Go语言中 case 是一个独立的代码块，执行完毕后不会像C语言那样紧接着执行下一个 case，
	//但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现这一功能，代码如下：
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}
