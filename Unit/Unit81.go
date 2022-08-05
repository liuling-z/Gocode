package main

import "fmt"

//Go语言链表操作
func main() {
	//单向链表
	//这里介绍三个概念：首元结点、头结点和头指针。
	//
	//    首元结点：就是链表中存储第一个元素的结点，如下图中 a1 的位置。
	//    头结点：它是在首元结点之前附设的一个结点，其指针域指向首元结点。头结点的数据域可以存储链表的长度或者其它的信息，也可以为空不存储任何信息。
	//    头指针：它是指向链表中第一个结点的指针。若链表中有头结点，则头指针指向头结点；若链表中没有头结点，则头指针指向首元结点。
	var head = new(Node)
	head.Data = 1
	var node1 = new(Node)
	node1.Data = 2
	head.Next = node1
	var node2 = new(Node)
	node2.Data = 3
	node1.Next = node2
	Shownode(head)

	head = new(Node)
	head.Data = 0
	var tail *Node
	tail = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = Node{Data: i}
		node.Next = tail //将新插入的node的next指向头结点
		tail = &node     //重新赋值头结点
	}
	Shownode(tail) //遍历结果
}

//使用 Struct 定义单链表
//【示例 1】使用 Struct 定义一个单向链表。
type Node struct {
	Data int
	Next *Node
}

func Shownode(p *Node) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
		fmt.Println(p)
		fmt.Println(*p)
		fmt.Println(*p)
	}
}
