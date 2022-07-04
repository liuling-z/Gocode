package main

import "fmt"

//Go语言结构体定义
func main() {
	//基本的实例化形式
	var LiRui People
	LiRui.name = "李锐"
	LiRui.age = "26"
	fmt.Println(LiRui)

	//创建指针类型的结构体
	Leo := new(People)
	Leo.name = "里奥"
	Leo.age = "35"
	Leo.wife.name = "伊丽莎白"
	Leo.child.name = "莉莉"
	fmt.Println(Leo)

	//取结构体的地址实例化
	Leo = &People{}
	Leo.name = "里奥"
	Leo.age = "35"
	Leo.wife.name = "伊丽莎白"
	Leo.wife.age = "35"
	Leo.child.name = "莉莉"
	Leo.child.age = "12"
	fmt.Println(Leo)
	type Command struct {
		Name    string // 指令名称
		Var     *int   // 指令绑定的变量
		Comment string // 指令的注释
	}
	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"
	fmt.Println(cmd.Comment)
	fmt.Println(cmd.Var)
	fmt.Println(*cmd)

	//代码说明如下：
	//
	//    第 1 行，定义 Command 结构体，表示命令行指令
	//    第 3 行，命令绑定的变量，使用整型指针绑定一个指针，指令的值可以与绑定的值随时保持同步。
	//    第 7 行，命令绑定的目标整型变量：版本号。
	//    第 9 行，对结构体取地址实例化。
	//    第 10～12 行，初始化成员字段。

	wife := new(Wife)
	wife.name = "阿尔托莉雅"
	child := new(Child)
	child.name = "莫德雷德"
	Leo = newPeople("里奥", "35", "男", *wife, *child)
	fmt.Println(Leo)

}

type People struct {
	name                 string
	age                  string
	gender               string
	address, city, phone string
	child                Child
	wife                 Wife
}

type Child struct {
	name   string
	age    string
	gender string
}
type Wife struct {
	name   string
	age    string
	gender string
}

func newPeople(name string, age string, gender string, wife2 Wife, child2 Child) *People {
	return &People{
		name:   name,
		age:    age,
		gender: gender,
		wife:   wife2,
		child:  child2,
	}
}
