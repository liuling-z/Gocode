package main

import "fmt"

//Go语言goto语句——跳转到指定的标签
func main() {
	//使用 goto 退出多层循环
	var breakAgain bool
	// 外循环
	for x := 0; x < 10; x++ {
		// 内循环
		for y := 0; y < 10; y++ {
			// 满足某个条件时, 退出循环
			if y == 2 {
				// 设置退出标记
				breakAgain = true
				// 退出本次循环
				break
			}
		}
		// 根据标记, 还需要退出一次循环
		if breakAgain {
			break
		}
	}
	fmt.Println("done")

}
