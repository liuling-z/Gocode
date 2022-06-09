package main

//Go语言sync.Map（在并发环境中使用的map）
func main() {
	//Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。
	//
	//下面来看下并发情况下读写 map 时会出现的问题，代码如下：

	// 创建一个int到int的映射
	m := make(map[int]int)
	// 开启一段并发代码
	go func() {
		// 不停地对map进行写入
		for {
			m[1] = 1
		}
	}()
	// 开启一段并发代码
	go func() {
		// 不停地对map进行读取
		for {
			_ = m[1]
		}
	}()
	// 无限循环, 让并发程序在后台执行
	for {
	}

}
