package main

import "fmt"

//Go语言初始化内嵌结构体
func main() {
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: Engine{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
}

// 车轮
type Wheel struct {
	Size int
}

// 引擎
type Engine struct {
	Power int    // 功率
	Type  string // 类型
}

// 车
type Car struct {
	Wheel
	Engine
}

//代码说明如下：
//
//    第 6 行定义车轮结构。
//    第 11 行定义引擎结构。
//    第 17 行定义车结构，由车轮和引擎结构体嵌入。
//    第 27 行，将 Car 的 Wheel 字段使用 Wheel 结构体进行初始化。
//    第 32 行，将 Car 的 Engine 字段使用 Engine 结构体进行初始化。
