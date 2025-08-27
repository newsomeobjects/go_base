package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// add1(i int) 是声明
// func() int 是返回类型
func add1(i int) func() int {
	return func() int {
		i++
		return i
	}
}

//goroutine 是轻量线程，
//
//但是 goroutine 不是通常理解的线程，线程是操作系统调度的。

//在 Go 中，想让某个任务并发或者异步执行，只需把任务封装为一个函数或闭包，交给 goroutine 执行即可。

func main() {

	go func() {
		fmt.Println("run goroutine in closure")
	}()

	go func(s string) {
		fmt.Println(s)
	}("gorouine: closure params")

	say("hello")

	count := add1(0) //这一步没有执行调用
	fmt.Println("-----------------", &count, count)
	fmt.Println(count()) //1
	fmt.Println(count()) //2
	fmt.Println(count())

}
