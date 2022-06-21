package demo

import (
	"fmt"
	"testing"
)

func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	} else if area == 2000 {
		fmt.Println("测试成功")
	}
}
