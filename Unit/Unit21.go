package main

import "fmt"

//遍历数组——访问每一个数组元素
func main() {
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	for k, v := range team {
		fmt.Println(k, v)
	}
	//	第 6 行，使用 for 循环，遍历 team 数组，遍历出的键 k 为数组的索引，值 v 为数组的每个元素值。
	//	第 7 行，将每个键值打印出来。

}
