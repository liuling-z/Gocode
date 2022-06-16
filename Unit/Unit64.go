package main

import (
	"fmt"
	"os"
	"sync"
)

//Go语言defer（延迟执行语句）
func main() {

	//多个延迟执行语句的处理顺序
	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")

}

// 使用延迟执行语句在函数退出时释放资源
//1) 使用延迟并发解锁
var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

// 根据键读取值
func readValue(key string) int {
	// 对共享资源加锁
	valueByKeyGuard.Lock()
	// 取值
	v := valueByKey[key]
	// 对共享资源解锁
	valueByKeyGuard.Unlock()
	// 返回值
	return v
}

//代码说明如下：
//
//    第 3 行，实例化一个 map，键是 string 类型，值为 int。
//    第 5 行，map 默认不是并发安全的，准备一个 sync.Mutex 互斥量保护 map 的访问。
//    第 9 行，readValue() 函数给定一个键，从 map 中获得值后返回，该函数会在并发环境中使用，需要保证并发安全。
//    第 11 行，使用互斥量加锁。
//    第 13 行，从 map 中获取值。
//    第 15 行，使用互斥量解锁。
//    第 17 行，返回获取到的 map 值。

// 使用 defer 语句对上面的语句进行简化，参考下面的代码。
func readValue2(key string) int {
	// 对共享资源加锁
	valueByKeyGuard.Lock()

	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

//上面的代码中第 6~8 行是对前面代码的修改和添加的代码，代码说明如下：
//
//    第 6 行在互斥量加锁后，使用 defer 语句添加解锁，该语句不会马上执行，而是等 readValue() 函数返回时才会被执行。
//    第 8 行，从 map 查询值并返回的过程中，与不使用互斥量的写法一样，对比上面的代码，这种写法更简单。

//2) 使用延迟释放文件句柄
// 根据文件名查询其大小
func fileSize(filename string) int64 {
	// 根据文件名打开文件, 返回文件句柄和错误
	f, err := os.Open(filename)
	// 如果打开时发生错误, 返回文件大小为0
	if err != nil {
		return 0
	}
	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误, 关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}
	// 取文件大小
	size := info.Size()
	// 关闭文件
	f.Close()

	// 返回文件大小
	return size
}

//代码说明如下：
//
//    第 2 行，定义获取文件大小的函数，返回值是 64 位的文件大小值。
//    第 5 行，使用 os 包提供的函数 Open()，根据给定的文件名打开一个文件，并返回操作文件用的句柄和操作错误。
//    第 8 行，如果打开的过程中发生错误，如文件没找到、文件被占用等，将返回文件大小为 0。
//    第 13 行，此时文件句柄 f 可以正常使用，使用 f 的方法 Stat() 来获取文件的信息，获取信息时，可能也会发生错误。
//    第 16～19 行对错误进行处理，此时文件是正常打开的，为了释放资源，必须要调用 f 的 Close() 方法来关闭文件，否则会发生资源泄露。
//    第 22 行，获取文件大小。
//    第 25 行，关闭文件、释放资源。
//    第 28 行，返回获取到的文件大小。

//在上面的例子中，第 25 行是对文件的关闭操作，下面使用 defer 对代码进行简化，代码如下：
func fileSize2(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	// 延迟调用Close, 此时Close不会被调用
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}
	size := info.Size()
	// defer机制触发, 调用Close关闭文件
	return size
}

//代码中加粗部分为对比前面代码而修改的部分，代码说明如下：
//
//    第 10 行，在文件正常打开后，使用 defer，将 f.Close() 延迟调用，注意，不能将这一句代码放在第 4 行空行处，一旦文件打开错误，f 将为空，在延迟语句触发时，将触发宕机错误。
//    第 16 行和第 22 行，defer 后的语句（f.Close()）将会在函数返回前被调用，自动释放资源。
