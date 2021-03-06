package main

import "fmt"

//Go语言从切片中删除元素
func main() {
	// 从开头位置删除
	//删除开头的元素可以直接移动数据指针：
	a := []int{1, 2, 3}
	a = a[1:] // 删除开头1个元素
	//也可以不移动数据指针，但是将后面的数据向开头移动，
	//可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）：
	a = []int{1, 2, 3}
	a = append(a[:0], a[1:]...) // 删除开头1个元素
	//还可以用 copy() 函数来删除开头的元素：
	a = []int{1, 2, 3}
	a = a[:copy(a, a[1:])] // 删除开头1个元素

	//从中间位置删除
	//对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用 append 或 copy 原地完成：
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a = append(a[:5], a[5+1:]...)  // 删除中间1个元素
	a = append(a[:5], a[5+3:]...)  // 删除中间3个元素
	a = a[:2+copy(a[2:], a[2+1:])] // 删除中间1个元素
	a = a[:2+copy(a[2:], a[2+5:])] // 删除中间5个元素

	//从尾部删除
	a = []int{1, 2, 3}
	a = a[:len(a)-1] // 删除尾部1个元素
	a = a[:len(a)-2] // 删除尾部2个元素

	//【示例】删除切片指定位置的元素。
	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
	//代码说明如下：
	//
	//    第 1 行，声明一个整型切片，保存含有从 a 到 e 的字符串。
	//    第 4 行，为了演示和讲解方便，使用 index 变量保存需要删除的元素位置。
	//    第 7 行，seq[:index] 表示的就是被删除元素的前半部分，值为 [1 2]，seq[index+1:] 表示的是被删除元素的后半部分，值为 [4 5]。
	//    第 10 行，使用 append() 函数将两个切片连接起来。
	//    第 12 行，输出连接好的新切片，此时，索引为 2 的元素已经被删除。
}
