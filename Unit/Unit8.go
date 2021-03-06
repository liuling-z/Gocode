package main

import "fmt"

/*
	\n：换行符
	\r：回车符
	\t：tab 键
	\u 或 \U：Unicode 字符
	\\：反斜杠自身

	字符串所占的字节长度可以通过函数 len() 来获取，例如 len(str)。

	字符串的内容（纯字节）可以通过标准索引法来获取，在方括号[ ]内写入索引，索引从 0 开始计数：
    字符串 str 的第 1 个字节：str[0]
    第 i 个字节：str[i - 1]
    最后 1 个字节：str[len(str)-1]

	字符串拼接符“+”

*/
func main() {
	var str = "C语言中文网\nGo语言教程"
	fmt.Println(len(str))
	fmt.Println(str)

	//定义多行字符串
	//	在Go语言中，使用双引号书写字符串的方式是字符串常见表达方式之一，被称为字符串字面量（string literal），这种双引号字面量不能跨行，如果想要在源码中嵌入一个多行字符串时，就必须使用`反引号，代码如下：
	//反引号`，是键盘上 1 键左边的键，两个反引号间的字符串将被原样赋值到 str 变量中。
	//在这种方式下，反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	str = `第一行\n
			第二行  \n
			第三行\r\n`
	fmt.Println(str)
}
