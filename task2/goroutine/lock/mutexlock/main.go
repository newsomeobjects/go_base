package main

import (
	"fmt"
	"sync"
)

/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/

func main() {
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	mx := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				func() {
					mx.Lock()
					count++
					defer mx.Unlock()     //保证解锁
					fmt.Println(i, count) //不同的协程在抢用
				}()

			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
