package main

import (
	"fmt"
	"math"
)

//整数、小数、复数、bool（布尔）类型
func main() {
	a := 3
	fmt.Printf("%d\n", a)
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	//复数
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))         // "10"

	//bool布尔
	var aVar = 10

	fmt.Println(aVar == 5)  // false
	fmt.Println(aVar == 10) // true
	fmt.Println(aVar != 5)  // true
	fmt.Println(aVar != 10) // false
}
