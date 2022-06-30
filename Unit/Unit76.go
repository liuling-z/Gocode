package main

import "fmt"

//Go语言构造函数
func main() {
	//其他编程语言构造函数的一些常见功能及特性如下：
	//
	//    每个类可以添加构造函数，多个构造函数使用函数重载实现。
	//    构造函数一般与类名同名，且没有返回值。
	//    构造函数有一个静态构造函数，一般用这个特性来调用父类的构造函数。
	//    对于 C++ 来说，还有默认构造函数、拷贝构造函数等。

	//多种方式创建和初始化结构体——模拟构造函数重载
	fmt.Println(NewCatByName("Tom").Name)
	fmt.Println(NewCatByColor("Black").Color)

	redCat := NewBlackCat("Red")
	fmt.Println(redCat)
	cat := NewCat("Tom")
	cat.Color = "Black"
	fmt.Println(cat)
}

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

//代码说明如下：
//
//    第 1 行定义 Cat 结构，包含颜色和名字字段。
//    第 6 行定义用名字构造猫结构的函数，返回 Cat 指针。
//    第 7 行取地址实例化猫的结构体。
//    第 8 行初始化猫的名字字段，忽略颜色字段。
//    第 12 行定义用颜色构造猫结构的函数，返回 Cat 指针。
//在这个例子中，颜色和名字两个属性的类型都是字符串，由于Go语言中没有函数重载，为了避免函数名字冲突，
//使用 NewCatByName() 和 NewCatByColor() 两个不同的函数名表示不同的 Cat 构造过程。

type BlackCat struct {
	Cat // 嵌入Cat, 类似于派生
}

// “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}
