package main

import (
	"fmt"
	"math"
)

//Go语言数据类型转换
func main() {
	z := 5.0
	x := int(z)
	fmt.Println(x)
	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	// 初始化一个32位整型值
	var a int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", a, a)
	// 将a变量数值转换为十六进制, 发生数值截断
	b := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", b, b)
	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))
	//代码说明如下：
	//
	//    第 11～14 行，输出几个常见整型类型的数值范围。
	//    第 17 行，声明 int32 类型的变量 a 并初始化。
	//    第 19 行，使用 fmt.Printf 的%x动词将数值以十六进制格式输出，这一行输出 a 在转换前的 32 位的值。
	//    第 22 行，将 a 的值转换为 int16 类型，也就是从 32 位有符号整型转换为 16 位有符号整型，由于 int16 类型的取值范围比 int32 类型的取值范围小，因此数值会进行截断（精度丢失）。
	//    第 24 行，输出转换后的 a 变量值，也就是 b 的值，同样以十六进制和十进制两种方式进行打印。
	//    第 27 行，math.Pi 是 math 包的常量，默认没有类型，会在引用到的地方自动根据实际类型进行推导，这里 math.Pi 被赋值到变量 c 中，因此类型为 float32。
	//    第 29 行，将 float32 转换为 int 类型并输出。
}
