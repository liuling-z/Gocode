package main

import "fmt"

//Go语言内嵌结构体成员名字冲突
func main() {
	c := &C{}
	c.A.a = 1
	fmt.Println(c)
}

type A struct {
	a int
}
type B struct {
	a int
}
type C struct {
	A
	B
}
