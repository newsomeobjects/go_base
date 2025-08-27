package main

import "fmt"

func TestLuoji() {
	a := true
	b := false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!(a && b))

	fmt.Println(0 & 0)
	fmt.Println(0 | 0)
	fmt.Println(0 ^ 0)
	fmt.Println()
	fmt.Println(0 & 1)
	fmt.Println(0 | 1)
	fmt.Println(0 ^ 1)
	fmt.Println()
	fmt.Println(1 & 1)
	fmt.Println(1 | 1)
	fmt.Println(1 ^ 1)
	fmt.Println()
	fmt.Println(1 & 0)
	fmt.Println(1 | 0)
	fmt.Println(1 ^ 0)

}
