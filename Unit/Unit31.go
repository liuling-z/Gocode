package main

import (
	"fmt"
	"sort"
	"strconv"
)

//Go语言map元素的删除和清空
func main() {
	//Go语言提供了一个内置函数 delete()，用于删除容器内的元素，
	//下面我们简单介绍一下如何用 delete() 函数删除 map 内的元素。
	//使用 delete() 函数从 map 中删除键值对
	//使用 delete() 内建函数从 map 中删除一组键值对，delete() 函数的格式如下：delete(map, 键)

	scene := make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Println(k, v)
	}

	//清空 map 中的所有元素
	//有意思的是，Go语言中并没有为 map 提供任何清空所有元素的函数、方法，
	//清空 map 的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，
	//Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
	scene = make(map[string]int)
	// 准备map数据
	for i := 0; i < 10; i++ {
		scene[strconv.Itoa(i)] = i
	}
	//注意：遍历输出元素的顺序与填充顺序无关，不能期望 map 在遍历时返回某种期望顺序的结果。
	//Java中此处会按顺序输出，但在Go语言中，如果需要特定顺序的遍历结果，正确的做法是先排序，代码如下：
	var sceneList []string
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	sort.Strings(sceneList)
	fmt.Println(sceneList)
}
