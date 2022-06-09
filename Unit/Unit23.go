package main

import "fmt"

//切片(slice)是对数组的一个连续片段的引用，
//所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，
//或者 Python 中的 list 类型），这个片段可以是整个数组，
//也可以是由起始和终止索引标识的一些项的子集，需要注意的是，
//终止索引标识的项不包括在切片内。
func main() {
	//从数组或切片生成新的切片
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
	//    从数组或切片生成新的切片拥有如下特性：
	//
	//    取出的元素数量为：结束位置 - 开始位置；
	//    取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
	//    当缺省开始位置时，表示从连续区域开头到结束位置；
	//    当缺省结束位置时，表示从开始位置到整个连续区域末尾；
	//    两者同时缺省时，与切片本身等效；
	//    两者同时为 0 时，等效于空切片，一般用于切片复位。
	//    根据索引位置取切片 slice 元素值时，取值范围是（0～len(slice)-1），
	//    超界会报运行时错误，生成切片时，结束位置可以填写 len(slice) 但不会报错。

	//1) 从指定范围中生成切片
	var highRiseBuilding [30]int
	for i := 0; i < len(highRiseBuilding); i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])

	//    代码说明如下：
	//
	//    第 8 行，尝试出租一个区间楼层。
	//    第 11 行，出租 20 层以上。
	//    第 14 行，出租 2 层以下，一般是商用铺面。

	//2) 表示原有的切片
	x := []int{1, 2, 3}
	fmt.Println(x[:])

	//3) 重置切片，清空拥有的元素
	y := []int{1, 2, 3}
	fmt.Println(y[0:0])

	//直接声明新的切片
	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)

	//    代码说明如下：
	//
	//    第 2 行，声明一个字符串切片，切片中拥有多个字符串。
	//    第 5 行，声明一个整型切片，切片中拥有多个整型数值。
	//    第 8 行，将 numListEmpty 声明为一个整型切片，本来会在{}中填充切片的初始化元素，这里没有填充，所以切片是空的，但是此时的 numListEmpty 已经被分配了内存，只是还没有元素。
	//    第 11 行，切片均没有任何元素，3 个切片输出元素内容均为空。
	//    第 14 行，没有对切片进行任何操作，strList 和 numList 没有指向任何数组或者其他切片。
	//    第 17 行和第 18 行，声明但未使用的切片的默认值是 nil，strList 和 numList 也是 nil，所以和 nil 比较的结果是 true。
	//    第 19 行，numListEmpty 已经被分配到了内存，但没有元素，因此和 nil 比较时是 false。

	//使用 make() 函数构造切片
	m := make([]int, 2)
	n := make([]int, 2, 10)
	fmt.Println(m, n)
	fmt.Println(len(m), len(n))

	//其中 m 和 n 均是预分配 2 个元素的切片，只是 n 的内部存储空间已经分配了 10 个，但实际使用了 2 个元素。
	//容量不会影响当前的元素个数，因此 m 和 n 取 len 都是 2。
}
