package main

import (
	"fmt"
)

//Go语言if else（分支结构）
func main() {
	//if condition1 {
	//	 do something
	//} else if condition2 {
	//	 do something else
	//}else {
	//	 catch-all or default
	//}

	var ten int = 11
	if ten > 10 {
		fmt.Println(">10")
	} else {
		fmt.Println("<=10")
	}

	//特殊写法
	if err := Connect(); err != nil {
		fmt.Println(err)
		return
	}
}

func Connect() map[string]string {
	scene := make(map[string]string)
	scene["001"] = "test"
	scene["002"] = "002"
	return scene
}
