package main

import (
	"errors"
	"fmt"
	"unsafe"
)

//Go语言nil：空值/零值
func main() {
	//在Go语言中，布尔类型的零值（初始值）为 false，
	//数值类型的零值为 0，字符串类型的零值为空字符串""，
	//而指针、切片、映射、通道、函数和接口的零值则是 nil。

	// nil 不是关键字或保留字
	//nil 并不是Go语言的关键字或者保留字，也就是说我们可以定义一个名称为 nil 的变量，比如下面这样：
	var nil = errors.New("my god")
	fmt.Println(nil)

	//nil 没有默认类型
	fmt.Printf("%T", nil)
	print(nil)

	//不同类型 nil 的指针是一样的
	var arr []int
	var num *int
	fmt.Println("\n不同类型 nil 的指针是一样的")
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p", num)

	//nil 是 map、slice、pointer、channel、func、interface 的零值
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	var p *struct{}
	var s []int
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)

	//不同类型的 nil 值占用的内存大小可能是不一样的
	fmt.Println(unsafe.Sizeof(p)) // 8
	fmt.Println(unsafe.Sizeof(s)) // 24
	fmt.Println(unsafe.Sizeof(m)) // 8
	fmt.Println(unsafe.Sizeof(c)) // 8
	fmt.Println(unsafe.Sizeof(f)) // 8
	fmt.Println(unsafe.Sizeof(i)) // 16
}
