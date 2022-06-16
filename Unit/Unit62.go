package main

import (
	"bytes"
	"fmt"
)

//获得可变参数类型——获得每一个参数的类型
func main() {
	// 将不同类型的变量通过printTypeValue()打印出来
	fmt.Println(printTypeValue(100, "str", true))
}
func printTypeValue(slist ...interface{}) string {
	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer
	// 遍历参数
	for _, s := range slist {
		// 将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)
		// 类型的字符串描述
		var typeString string
		// 对s进行类型断言
		switch s.(type) {
		case bool: // 当s为布尔类型时
			typeString = "bool"
		case string: // 当s为字符串类型时
			typeString = "string"
		case int: // 当s为整型类型时
			typeString = "int"
		}
		// 写字符串前缀
		b.WriteString("value: ")
		// 写入值
		b.WriteString(str)
		// 写类型前缀
		b.WriteString(" type: ")
		// 写类型字符串
		b.WriteString(typeString)
		// 写入换行符
		b.WriteString("\n")
	}
	return b.String()
}

//代码说明如下：
//
//    第 8 行，printTypeValue() 输入不同类型的值并输出类型和值描述。
//    第 11 行，bytes.Buffer 字节缓冲作为快速字符串连接。
//    第 14 行，遍历 slist 的每一个元素，类型为 interface{}。
//    第 17 行，使用 fmt.Sprintf 配合%v动词，可以将 interface{} 格式的任意值转为字符串。
//    第 20 行，声明一个字符串，作为变量的类型名。
//    第 23 行，switch s.(type) 可以对 interface{} 类型进行类型断言，也就是判断变量的实际类型。
//    第 24～29 行为 s 变量可能的类型，将每种类型的对应类型字符串赋值到 typeString 中。
//    第 33～42 行为写输出格式的过程。
