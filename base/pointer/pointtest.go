package main

import "fmt"

var p1 *int

var p2 *string

var t1 *int
var t2 *string

var pp **int

var pp1 *******int

func main() {
	i := 1
	s := "Hello"
	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
	// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	p1 = &i
	p2 = &s

	p3 := &p2

	fmt.Println(p1) // i的地址

	fmt.Println(p2) // s的地址

	fmt.Println(p3) // p2的地址

	//指针默认值为nil
	fmt.Println(t1)
	fmt.Println(t2)
	//打印声明的 指针 的地址
	fmt.Println(&t1)

	s1 := "asdf"
	var o1 *string = &s1

	var o2 **string = &o1

	var o3 ***string = &o2

	o22 := &o1

	fmt.Println("---------------")
	fmt.Println(s1)
	fmt.Println(*o1)
	fmt.Println(**o2)
	fmt.Println(**o22)

	fmt.Println("-------o3 -> *o3 -> **o3 -> ***03--------")
	fmt.Println(o3)
	fmt.Println(*o3)
	fmt.Println(**o3)
	fmt.Println(***o3)

	fmt.Println("-------o3 -- o2 -- o1 -- s1--------")
	fmt.Println(o3)
	fmt.Println(o2)
	fmt.Println(o1)
	fmt.Println(s1)

	fmt.Println("===update===")
	***o3 = "yyyy"
	fmt.Println(s1)

	fmt.Println("-------&o3 -- &o2 -- &o1 -- &s1--------")
	fmt.Println(&o3)
	fmt.Println(&o2)
	fmt.Println(&o1)
	fmt.Println(&s1)

}
