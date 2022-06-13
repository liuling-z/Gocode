package main

import (
	"container/list"
	"fmt"
)

//遍历列表——访问列表的每一个元素
func main() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	//正向遍历
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	//反向遍历
	for i := l.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
	//代码说明如下：
	//
	//    第 1 行，创建一个列表实例。
	//    第 4 行，将 canon 放入列表尾部。
	//    第 7 行，在队列头部放入 67。
	//    第 9 行，使用 for 语句进行遍历，其中 i:=l.Front() 表示初始赋值，只会在一开始执行一次，
	//    每次循环会进行一次 i != nil 语句判断，如果返回 false，表示退出循环，反之则会执行 i = i.Next()。
	//    第 10 行，使用遍历返回的 *list.Element 的 Value 成员取得放入列表时的原值。
}
