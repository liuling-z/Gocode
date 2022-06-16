package main

import "fmt"

//Go语言函数类型实现接口——把函数作为接口来调用
func main() {
	// 声明接口变量
	var invoker Invoker
	// 实例化结构体
	s := new(Struct)
	// 将实例化的结构体赋值到接口
	invoker = s
	// 使用接口调用实例化结构体的方法Struct.Call
	invoker.Call("hello")
	// 将匿名函数转为FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	// 使用接口调用FuncCaller.Call，内部会调用函数本体
	invoker.Call("hello")

	//代码说明如下：
	//
	//    第 2 行，声明 Invoker 类型的变量。
	//    第 5 行，使用 new 将结构体实例化，此行也可以写为 s:=&Struct。
	//    第 8 行，s 类型为 *Struct，已经实现了 Invoker 接口类型，因此赋值给 invoker 时是成功的。
	//    第 11 行，通过接口的 Call() 方法，传入 hello，此时将调用 Struct 结构体的 Call() 方法。
	//代码说明如下：
	//
	//    第 2 行，声明接口变量。
	//    第 5 行，将 func(v interface{}){} 匿名函数转换为 FuncCaller 类型（函数签名才能转换），此时 FuncCaller 类型实现了 Invoker 的 Call() 方法，赋值给 invoker 接口是成功的。
	//    第 10 行，使用接口方法调用。
}

// 调用器接口
type Invoker interface {
	// 需要实现一个Call方法
	Call(interface{})
}

// 结构体类型
type Struct struct {
}

// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

//代码说明如下：
//
//    第 2 行，定义结构体，该例子中的结构体无须任何成员，主要展示实现 Invoker 的方法。
//    第 6 行，Call() 为结构体的方法，该方法的功能是打印 from struct 和传入的 interface{} 类型的值。

// 函数定义为类型
type FuncCaller func(interface{})

// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	// 调用f函数本体
	f(p)
}

//代码说明如下：
//
//    第 2 行，将 func(interface{}) 定义为 FuncCaller 类型。
//    第 5 行，FuncCaller 的 Call() 方法将实现 Invoker 的 Call() 方法。
//    第 8 行，FuncCaller 的 Call() 方法被调用与 func(interface{}) 无关，还需要手动调用函数本体。
