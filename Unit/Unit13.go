package main

import (
	"flag"
	"fmt"
)

//示例：使用指针变量获取命令行的输入信息

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main() {
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)
	fmt.Println(mode)

	//	代码说明如下：
	//
	//    第 10 行，通过 flag.String，定义一个 mode 变量，这个变量的类型是 *string。后面 3 个参数分别如下：
	//        参数名称：在命令行输入参数时，使用这个名称。
	//        参数值的默认值：与 flag 所使用的函数创建变量类型对应，String 对应字符串、Int 对应整型、Bool 对应布尔型等。
	//        参数说明：使用 -help 时，会出现在说明中。
	//    第 15 行，解析命令行参数，并将结果写入到变量 mode 中。
	//    第 18 行，打印 mode 指针所指向的变量。
}
