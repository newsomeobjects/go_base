package main

import "fmt"

func TestBool() {
	a := 1
	b := 5

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	//只能同类型比较
	//var c int64 = 1
	//fmt.Println(a <= c)
}
