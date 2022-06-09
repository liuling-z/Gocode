package main

import (
	"fmt"
	"reflect"
	"time"
)

//非本地类型不能定义方法
//在结构体成员嵌入时使用别名
func main() {
	// 声明变量a为车辆类型
	var a Vehicle

	// 指定调用FakeBrand的Show
	a.FakeBrand.Show()
	// 取a的类型反射对象
	ta := reflect.TypeOf(a)
	// 遍历a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.
			Name())
	}
}

// 定义time.Duration的别名为MyDuration
type MyDuration time.Duration

// 为MyDuration添加一个函数
func (m MyDuration) EasySet(a string) {
}

// 定义商标结构
type Brand struct {
}

// 为商标结构添加Show()方法
func (t Brand) Show() {
}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

//代码说明如下：
//
//    第 9 行，定义商标结构。
//    第 13 行，为商标结构添加 Show() 方法。
//    第 17 行，为 Brand 定义一个别名 FakeBrand。
//    第 20～25 行，定义车辆结构 Vehicle，嵌入 FakeBrand 和 Brand 结构。
//    第 30 行，将 Vechicle 实例化为 a。
//    第 33 行，显式调用 Vehicle 中 FakeBrand 的 Show() 方法。
//    第 36 行，使用反射取变量 a 的反射类型对象，以查看其成员类型。
//    第 39～42 行，遍历 a 的结构体成员。
//    第 45 行，打印 Vehicle 类型所有成员的信息。
