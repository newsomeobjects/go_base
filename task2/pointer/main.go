package main

import "fmt"

/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func modify(a *int) {
	*a += 10
}
func modify1(a int) {
	a += 10
}

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func modifySlice(slice *[]int) {
	s := *slice
	for i := range s {
		//让每个元素乘二
		s[i] *= 2
	}
}

func modifySlice1(slice []int) {
	for i := range slice {
		slice[i] *= 2
	}
}

func modifyarray(a *[2]int) {
	a[0] *= 2
}

func append1(slice []int) {
	slice = append(slice, 1, 2, 2, 2, 2, 2, 2, 2) //引起扩容，底层数组指针被修改，所以操作不会影响到原值
}

func main() {

	a := 5
	modify(&a)
	fmt.Println(a) //15
	modify1(a)
	fmt.Println(a) //15

	b := [2]int{2, 3}
	modifyarray(&b)
	fmt.Println(b)

	s := []int{3, 4, 5, 1, 23, 54}
	modifySlice(&s)
	fmt.Println(s)
	modifySlice1(s)
	fmt.Println(s)
	append1(s)
	fmt.Println("append:", s)

	//切片无法直接拷贝,数组可以这样做
	ss := s
	ss[0] = 1
	fmt.Println("ss:", ss)
	fmt.Println("s:", s)

	sss := s[:]
	sss[0] = 2
	fmt.Println("sss:", sss)
	fmt.Println("s:", s)

}

/*
Go语言中所有的函数传参都是值传递（Pass by Value）

引用类型：
切片（Slices）
映射（Maps）
通道（Channels）
函数（Functions）
接口（Interfaces）
它们是一种内部已经包含指针的、复杂的底层数据结构。
可能导致原值没有修改的错误：
切片append 操作导致扩容后底层数组的指针被修改
map被重新赋值后，指针被修改
接口被重新赋值后，指针被修改

值类型：
基本类型：如 int, float64, bool, string, rune 等。
数组（Arrays）：注意，数组和切片不同。数组是固定长度的值类型。
结构体（Structs）：尤其是当你需要修改结构体字段或在函数间传递大结构体以避免拷贝开销时。
自定义类型（Type Definitions）：基于上述类型定义的新类型，例如 type MyInt int。
---你必须传递这个原值的地址（即指针），让函数通过地址找到原始内存位置并进行修改。



拷贝：
值类型 默认就是深拷贝， 直接拷贝，直接赋值即可拷贝，
基本数据类型、数组、结构体

引用类型：深拷贝
切片 copy() 函数或 append
map 手动遍历拷贝
其他的引用类型：可以序列化，递归拷贝。。

切片类型浅拷贝
slice2 := slice1	浅拷贝	最直接的浅拷贝方式
slice2 := slice1[:]	浅拷贝	使用切片表达式
函数参数传递	浅拷贝	传递切片头副本
共享底层数组,容量共享


copy(slice2, slice1)	深拷贝	创建独立副本
slice2 := append([]int(nil), slice1...) // 深拷贝
*/
