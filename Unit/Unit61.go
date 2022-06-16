package main

import (
	"bytes"
	"fmt"
)

//遍历可变参数列表——获取每一个参数的值
func main() {
	// 输入3个字符串, 将它们连成一个字符串
	fmt.Println(joinStrings("pig ", "and", " rat"))
	fmt.Println(joinStrings("hammer", " mom", " and", " hawk"))
}
func joinStrings(slist ...string) string {
	// 定义一个字节缓冲, 快速地连接字符串
	var b bytes.Buffer
	// 遍历可变参数列表slist, 类型为[]string
	for _, s := range slist {
		// 将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}
	// 将连接好的字节数组转换为字符串并输出
	return b.String()
}

//代码说明如下：
//
//    第 8 行，定义了一个可变参数的函数，slist 的类型为 []string，每一个参数的类型都是 string，也就是说，该函数只接受字符串类型作为参数。
//    第 11 行，bytes.Buffer 在这个例子中的作用类似于 StringBuilder，可以高效地进行字符串连接操作。
//    第 13 行，遍历 slist 可变参数，s 为每个参数的值，类型为 string。
//    第 15 行，将每一个传入参数放到 bytes.Buffer 中。
//    第 19 行，将 bytes.Buffer 中的数据转换为字符串作为函数返回值返回。
//    第 24 行，输入 3 个字符串，使用 joinStrings() 函数将参数连接为字符串输出。
//    第 25 行，输入 4 个字符串，连接后输出。
