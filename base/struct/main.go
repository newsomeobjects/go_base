package main

import (
	"fmt"
	"go_base/base/struct/a"
	"go_base/base/struct/b"
)

/*
最小暴露原则：只公开必要的方法

保持一致性：同一个结构体的方法使用统一的接收者类型

方法分组：通常一个结构体会有：

几个公开的"接口方法"

多个私有的"辅助方法"
*/
func main() {
	a1 := a.NewA()

	fmt.Println(a1)

	b1 := b.NewB()
	fmt.Println(b1)

	//copy deep
	c := b1
	c.SetB("steadfast")
	fmt.Println(c)

	d := &b1
	d.SetB("通过引用地址修改变量，打印b1看看其参数修改了吗")
	fmt.Println(d) //打印结构体地址，发现打印的是值，加了&前缀
	fmt.Println(*d)

	//引用类型的零值是 nil
	//
	//值类型的零值是各自的默认值（0、false、""等）
	//
	//接口变量只有在没有存储任何具体值时才是 nil
	var ee a.A
	fmt.Println(ee)
	var ee1 b.B
	fmt.Println(ee1)
}
