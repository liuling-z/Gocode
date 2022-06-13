package main

import (
	"container/list"
	"fmt"
)

//Go语言list（列表）
func main() {
	//列表是一种非连续的存储容器，由多个节点组成，
	//节点通过一些变量记录彼此之间的关系，列表有多种实现方法，如单链表、双链表等。
	//
	//初始化列表
	//list 的初始化有两种方法：分别是使用 New() 函数和 var 关键字声明，两种方法的初始化效果都是一致的。
	//1) 通过 container/list 包的 New() 函数初始化 list
	//变量名 := list.New()
	//2) 通过 var 关键字声明初始化 list
	//var 变量名 list.List
	//列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型，这既带来了便利，也引来一些问题，
	//例如给列表中放入了一个 interface{} 类型的值，取出值后，如果要将 interface{} 转换为其他类型将会发生宕机。

	//在列表中插入元素
	//双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
	// 提示
	//这两个方法都会返回一个 *list.Element 结构，如果在以后的使用中需要删除插入的元素，
	//则只能通过 *list.Element 配合 Remove() 方法进行删除，这种方法可以让删除更加效率化，同时也是双链表特性之一。

	l := list.New()
	fmt.Println(l.PushBack("fist"))
	fmt.Println(l.PushFront(67))
	fmt.Println(l.PushFront("test"))
	fmt.Println(l.Front())
	fmt.Println(l.Back())

	//代码说明如下：
	//
	//    第 1 行，创建一个列表实例。
	//    第 3 行，将 fist 字符串插入到列表的尾部，此时列表是空的，插入后只有一个元素。
	//    第 4 行，将数值 67 放入列表，此时，列表中已经存在 fist 元素，67 这个元素将被放在 fist 的前面。

	//列表插入元素的方法如下表所示。
	//
	//方  法 															功  能
	//InsertAfter(v interface {}, mark * Element) * Element 	在 mark 点之后插入元素，mark 点由其他插入函数提供
	//InsertBefore(v interface {}, mark * Element) *Element 	在 mark 点之前插入元素，mark 点由其他插入函数提供
	//PushBackList(other *List) 								添加 other 列表元素到尾部
	//PushFrontList(other *List) 								添加 other 列表元素到头部
}
