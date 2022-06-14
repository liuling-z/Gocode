package main

import "fmt"

//Go语言for（循环结构）
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println(sum)

JLoop:
	for j := 0; j < 5; j++ {
	SLoop:
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
				break SLoop
			}
			fmt.Println(i)
		}
	}

}
