package main

import "fmt"

//Go语言类型内嵌和结构体内嵌
func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Printf("outer2 is:", outer2)

	//内嵌结构体
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println("\n", b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)

	//结构内嵌特性
	// 1) 内嵌的结构体可以直接访问其成员变量
	//嵌入结构体的成员，可以通过外部结构体的实例直接访问。如果结构体有多层嵌入结构体，结构体实例访问任意一级的嵌入结构体成员时都只用给出字段名，
	//而无须像传统结构体字段一样，通过一层层的结构体字段访问到最终的字段。例如，ins.a.b.c的访问可以简化为ins.c。

	// 2) 内嵌结构体的字段名是它的类型名
	//内嵌结构体字段仍然可以使用详细的字段进行一层层访问，内嵌结构体的字段名就是它的类型名，代码如下：
	//var c Color
	//c.BasicColor.R = 1
	//c.BasicColor.G = 1
	//c.BasicColor.B = 0

}

type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}
type A struct {
	ax, ay int
}
type B struct {
	A
	bx, by float32
}
