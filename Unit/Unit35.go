package main

import "container/list"

//从列表(List)中删除元素
func main() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)

	//代码说明如下：
	//第 6 行，创建列表实例。
	//第 9 行，将字符串 canon 插入到列表的尾部。
	//第 12 行，将数值 67 添加到列表的头部。
	//第 15 行，将字符串 fist 插入到列表的尾部，并将这个元素的内部结构保存到 element 变量中。
	//第 18 行，使用 element 变量，在 element 的位置后面插入 high 字符串。
	//第 21 行，使用 element 变量，在 element 的位置前面插入 noon 字符串。
	//第 24 行，移除 element 变量对应的元素。

	//操作内容 							列表元素
	//l.PushBack("canon") 				canon
	//l.PushFront(67) 					67, canon
	//element := l.PushBack("fist") 	67, canon, fist
	//l.InsertAfter("high", element) 	67, canon, fist, high
	//l.InsertBefore("noon", element) 	67, canon, noon, fist, high
	//l.Remove(element) 				67, canon, noon, high
}
