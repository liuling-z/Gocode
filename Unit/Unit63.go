package main

import "fmt"

//在多个可变参数函数中传递参数
func main() {
	print(1, 2, 3)
}

// 实际打印的函数
func rawPrint(rawList ...interface{}) {
	// 遍历可变参数切片
	for _, a := range rawList {
		// 打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
func print(slist ...interface{}) {
	// 将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)
}

//对代码的说明：
//
//    第 9～13 行，遍历 rawPrint() 的参数列表 rawList 并打印。
//    第 20 行，将变量在 print 的可变参数列表中添加...后传递给 rawPrint()。
//    第 25 行，传入 1、2、3 这 3 个整型值并进行打印。
