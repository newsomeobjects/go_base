package main

import "fmt"

// 方式1
const a int = 1

// 方式2
const b = "test"

// 方式3
const c, d = 2, "hello"

// 方式4
const e, f bool = true, false

// 方式5
const (
	h    byte = 3
	i         = "value"
	j, k      = "v", 4
	l, m      = 5, false
)

const (
	n = 6
)

// Go 中，**常量只能使用基本数据类型**，即数字、字符串和布尔类型。
// 不能使用复杂的数据结构，比如切片、数组、map、指针和结构体等。
// 如果使用了非基本数据类型，会在编译期报错。
func main() {

	a := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := a[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := a[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := a[:2]
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
	a[0] = 9
	a[1] = 8
	a[2] = 7
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)

	switch d := 1; d {
	case 1:
		e := 4
		fmt.Println("declare e = ", e)
		fmt.Println("d == 1")
	case 3:
		f := 4
		fmt.Println("declare f = ", f)
		fmt.Println("d == 3")
	}

	aar := [5]int{1, 2, 3, 4}
	fmt.Println(aar)
	aar = [5]int{3: 1, 0: 2, 5}
	fmt.Println(aar)
}
