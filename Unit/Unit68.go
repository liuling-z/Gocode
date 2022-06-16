package main

import (
	"errors"
	"fmt"
)

//Go语言处理运行时错误
func main() {
	//Go语言的错误处理思想及设计包含以下特征：
	//
	//    一个可能造成错误的函数，需要返回值中返回一个错误接口（error），如果调用是成功的，错误接口将返回 nil，否则返回错误。
	//    在函数调用后需要检查错误，如果发生错误，则进行必要的错误处理。

	//错误接口的定义格式
	type error interface {
		Error() string
	}

	//自定义一个错误
	var err = errors.New("this is an error")
	fmt.Println(err)

	fmt.Println(div(1, 0))

}

//1) errors 包
// 创建错误对象
func New(text string) error {
	return &errorString{text}
}

// 错误字符串
type errorString struct {
	s string
}

// 返回发生何种错误
func (e *errorString) Error() string {
	return e.s
}

//代码说明如下：
//
//    第 2 行，将 errorString 结构体实例化，并赋值错误描述的成员。
//    第 7 行，声明 errorString 结构体，拥有一个成员，描述错误内容。
//    第 12 行，实现 error 接口的 Error() 方法，该方法返回成员中的错误描述。

//2) 在代码中使用错误定义
// 定义除数为0的错误
var errDivisionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {
	// 判断除数为0的情况并返回
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	// 正常计算，返回空错误
	return dividend / divisor, nil
}

//代码说明：
//
//    第 9 行，预定义除数为 0 的错误。
//    第 11 行，声明除法函数，输入被除数和除数，返回商和错误。
//    第 14 行，在除法计算中，如果除数为 0，计算结果为无穷大，为了避免这种情况，对除数进行判断，并返回商为 0 和除数为 0 的错误对象。
//    第 19 行，进行正常的除法计算，没有发生错误时，错误对象返回 nil。
