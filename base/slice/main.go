package main

import "fmt"

func main() {
	var a []int
	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	fmt.Println(a)

	a = append(a[:1], append([]int{8}, a[1:]...)...) // 在第1个位置插入8
	fmt.Println(a)

	//a = append(a[:1], append([]int{1, 2, 3}, a[1:]...)...) // 在第i个位置插入切片

	//fmt.Println(a)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice1, "  ", slice2)

	slice1 = []int{1, 2, 3, 4, 5}
	slice2 = []int{5, 4, 3}
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1, "  ", slice2)

	/*
	   	Go语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素，
	   根据要删除元素的位置有三种情况，
	   分别是从开头位置删除、从中间位置删除和从尾部删除，其中删除切片尾部的元素速度最快。
	*/

}
