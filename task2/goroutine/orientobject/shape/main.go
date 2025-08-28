package main

import "fmt"

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (rec *Rectangle) Perimeter() {
	fmt.Println("Perimeter")
}

func (rec *Rectangle) Area() {
	fmt.Println("area")
}

type Circle struct {
}

func (c Circle) Area() {
	fmt.Println("area")
}

func (c Circle) Perimeter() {
	fmt.Println("Perimeter")
}

func main() {
	c := Circle{}
	c.Perimeter()
	c.Area()
	var ss Shape = c
	fmt.Println(ss)

	d := Rectangle{}
	var sss Shape = &d //必须传地址，因为是接口实现接收者用的是指针
	fmt.Println(sss)
} //要实现接口里所有方法
