package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func main() {
	var count int64
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i1 := 0; i1 < 10; i1++ {
		go func(j int) { //---------------------必须值传递进去for循环的额变量
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&count, 1)
				fmt.Println(i1, count)
			}
		}(i1)
	}

	wg.Wait()
	fmt.Println(count)
}
