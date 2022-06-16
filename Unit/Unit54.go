package main

import (
	"flag"
	"fmt"
)

//使用匿名函数实现操作封装
var skillParam = flag.String("skill", "fire", "skill to perform")

func main() {
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
	//代码说明如下：
	//
	//    第 8 行，定义命令行参数 skill，从命令行输入 --skill 可以将 = 后的字符串传入 skillParam 指针变量。
	//    第 12 行，解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值。
	//    第 14 行，定义一个从字符串映射到 func() 的 map，然后填充这个 map。
	//    第 15～23 行，初始化 map 的键值对，值为匿名函数。
	//    第 26 行，skillParam 是一个 *string 类型的指针变量，使用 *skillParam 获取到命令行传过来的值，并在 map 中查找对应命令行参数指定的字符串的函数。
	//    第 29 行，如果在 map 定义中存在这个参数就调用，否则打印“技能没有找到”。
}
