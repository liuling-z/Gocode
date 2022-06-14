package main

import "fmt"

//Go语言break（跳出循环）
func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
	//代码说明如下：
	//
	//    第 7 行，外层循环的标签。
	//    第 8 行和第 9 行，双层循环。
	//    第 10 行，使用 switch 进行数值分支判断。
	//    第 13 和第 16 行，退出 OuterLoop 对应的循环之外，也就是跳转到第 20 行。
}
