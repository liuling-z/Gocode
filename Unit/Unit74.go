package main

import "fmt"

//Go语言初始化结构体的成员变量
func main() {
	//使用“键值对”初始化结构体
	//1) 键值对初始化结构体的书写格式
	week := &week{
		Monday:    "周一",
		Tuesday:   "周二",
		Wednesday: "周三",
		Thursday:  "周四",
		Friday:    "周五",
		Saturday:  "周六",
		Sunday:    "周日",
	}
	fmt.Println(week)
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr)

	//初始化匿名结构体
	//1) 匿名结构体定义格式和初始化写法
	cat := struct {
		name string
	}{name: "Tom"}

	fmt.Println(cat)
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

type week struct {
	Monday    string
	Tuesday   string
	Wednesday string
	Thursday  string
	Friday    string
	Saturday  string
	Sunday    string
}
