package main

import "fmt"

type SameFieldA struct {
	name  string
	value int
}

type SameFieldB struct {
	name  string
	value int
}

func (s *SameFieldB) getValue() int {
	return s.value
}

func main() {
	a := SameFieldA{
		name:  "a",
		value: 1,
	}

	b := SameFieldB(a)
	//结构体名称、类型得一致 ，数量也得一致

	fmt.Printf("conver SameFieldA to SameFieldB, value is : %d \n", b.getValue())

	// 只能结构体类型实例之间相互转换，指针不可以相互转换
	// var c interface{} = &a
	// _, ok := c.(*SameFieldB)
	// fmt.Printf("c is *SameFieldB: %v \n", ok)
}
