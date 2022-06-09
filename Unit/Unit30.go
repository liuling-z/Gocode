package main

import (
	"fmt"
	"sort"
)

//Go语言遍历map（访问map中的每一个键值对）
func main() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
	}
	//遍历对于Go语言的很多对象来说都是差不多的，直接使用 for range 语法即可，
	//遍历时，可以同时获得键和值，如只遍历值，可以使用下面的形式：for _, v := range scene {
	//将不需要的键使用_改为匿名变量形式。

	//只遍历键时，无须将值改为匿名变量形式，忽略值即可。使用下面的形式：
	//for k := range scene {

	//如果需要特定顺序的遍历结果，正确的做法是先排序，代码如下：
	scene = make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	// 声明一个切片保存map数据
	var sceneList []string
	// 将map数据遍历复制到切片中
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	// 对切片进行排序
	sort.Strings(sceneList)
	// 输出
	fmt.Println(sceneList)

	//代码说明如下：
	//
	//    第 1 行，创建一个 map 实例，键为字符串，值为整型。
	//    第 4～6 行，将 3 个键值对写入 map 中。
	//    第 9 行，声明 sceneList 为字符串切片，以缓冲和排序 map 中的所有元素。
	//    第 12 行，将 map 中元素的键遍历出来，并放入切片中。
	//    第 17 行，对 sceneList 字符串切片进行排序，排序时，sceneList 会被修改。
	//    第 20 行，输出排好序的 map 的键。
}
