package main

import "fmt"

func TestNumber() {
	//算术运算符
	//当不同的数字类型混合计算时，必须先把它们转换成同一类型才可以计算：

	a := 10 + 0.1
	b := byte(1) + 1
	fmt.Println(a, b)

	sum := a + float64(b)
	fmt.Println(sum)

	sub := byte(a) - b
	fmt.Println(sub)

	mul := a * float64(b)
	div := byte(a) / b

	fmt.Println(mul, div)
}
