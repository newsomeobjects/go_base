package main

import "fmt"

func main() {
	fmt.Println("-----------算术--------------")
	TestNumber()
	fmt.Println("-----------关系--------------")
	TestBool()

	fmt.Println("-----------逻辑---------------")
	TestLuoji()

	fmt.Println("-----------赋值---------------")
	TestFuZhi()

	fmt.Println("-----------其他---------------")
	a := 4
	var ptr *int
	fmt.Println(a)
	ptr = &a
	fmt.Printf("*ptr 为 %d\n", *ptr)

	fmt.Println("-----------优先级---------------")
	TestYouXian()
}
