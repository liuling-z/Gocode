package main

import "fmt"

//使用 goto 退出多层循环
func main() {
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
	//代码说明如下：
	//
	//    第 10 行，构建外循环。
	//    第 13 行，构建内循环。
	//    第 16 行，当 y==2 时需要退出所有的 for 循环。
	//    第 19 行，默认情况下循环只能一层一层退出，为此就需要设置一个状态变量 breakAgain，需要退出时，设置这个变量为 true。
	//    第 22 行，使用 break 退出当前循环，执行后，代码调转到第 28 行。
	//    第 28 行，退出一层循环后，根据 breakAgain 变量判断是否需要再次退出外层循环。
	//    第 34 行，退出所有循环后，打印 done。

	//将上面的代码使用Go语言的 goto 语句进行优化：
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")

	//代码说明如下：
	//
	//    第 13 行，使用 goto 语句跳转到指明的标签处，标签在第 23 行定义。
	//    第 20 行，标签只能被 goto 使用，但不影响代码执行流程，此处如果不手动返回，在不满足条件时，也会执行第 24 行代码。
	//    第 23 行，定义 breakHere 标签。
	//使用 goto 语句后，无须额外的变量就可以快速退出所有的循环。

	//使用 goto 集中处理错误
	//	err := firstCheckError()
	//	if err != nil {
	//		goto onExit
	//	}
	//	err = secondCheckError()
	//	if err != nil {
	//		goto onExit
	//	}
	//	fmt.Println("done")
	//	return
	//onExit:
	//	fmt.Println(err)
	//	exitProcess()
}
