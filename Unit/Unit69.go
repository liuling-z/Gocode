package main

import "fmt"

//示例：在解析中使用自定义错误
func main() {
	var e error
	// 创建一个错误实例，包含文件名和行号
	e = newParseError("BaiduRetriever.go", 1)
	// 通过error接口查看错误描述
	fmt.Println(e.Error())
	// 根据错误接口具体的类型，获取详细错误信息
	switch detail := e.(type) {
	case *ParseError: // 这是一个解析错误
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default: // 其他类型的错误
		fmt.Println("other error")
	}
}

// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

//代码说明如下：
//
//    第 8 行，声明了一个解析错误的结构体，解析错误包含有 2 个成员，分别是文件名和行号。
//    第 14 行，实现了错误接口，将成员的文件名和行号格式化为字符串返回。
//    第 19 行，根据给定的文件名和行号创建一个错误实例。
//    第 25 行，声明一个错误接口类型。
//    第 27 行，创建一个实例，这个错误接口内部是 *ParserError 类型，携带有文件名 BaiduRetriever.go 和行号 1。
//    第 30 行，调用 Error() 方法，通过第 15 行返回错误的详细信息。
//    第 33 行，通过错误断言，取出发生错误的详细类型。
//    第 34 行，通过分析这个错误的类型，得知错误类型为 *ParserError，此时可以获取到详细的错误信息。
//    第 36 行，如果不是我们能够处理的错误类型，会打印出其他错误做出其他的处理。
