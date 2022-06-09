package main

import "fmt"

//Go语言map（Go语言映射）
func main() {
	//Go语言中 map 是一种特殊的数据结构，一种元素对（pair）的无序集合，
	//pair 对应一个 key（索引）和一个 value（值），所以这个结构也称为关联数组或字典，
	//这是一种能够快速寻找值的理想结构，给定 key，就可以迅速找到对应的 value。
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	//示例中 mapLit 演示了使用{key1: value1, key2: value2}的格式来初始化 map ，就像数组和结构体一样。
	//
	//上面代码中的 mapCreated 的创建方式mapCreated := make(map[string]float)等价于mapCreated := map[string]float{} 。
	//
	//mapAssigned 是 mapList 的引用，对 mapAssigned 的修改也会影响到 mapLit 的值。

	//map 容量
	//和数组不同，map 可以根据新增的 key-value 动态的伸缩，
	//因此它不存在固定长度或者最大限制，但是也可以选择标明 map 的初始容量 capacity，格式如下：
	noteFrequency := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
	fmt.Println(noteFrequency)

	//用切片作为 map 的值
	//既然一个 key 只能对应一个 value，而 value 又是一个原始类型，
	//那么如果一个 key 要对应多个值怎么办？例如，当我们要处理 unix 机器上的所有进程，
	//以父进程（pid 为整形）作为 key，所有的子进程（以所有子进程的 pid 组成的切片）作为 value。
	//通过将 value 定义为 []int 类型或者其他类型的切片，就可以优雅的解决这个问题，示例代码如下所示：
	mp1 := make(map[int][]int)
	mp2 := make(map[int]*[]int)
	fmt.Println(mp1)
	fmt.Println(mp2)
}
