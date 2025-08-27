package main

import "fmt"

func main() {
	// 九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, i*j)
		}
		fmt.Println()
	}

	//range
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	m := map[string]int{
		"hello": 100,
		"world": 200,
	}

	//对 map 遍历时，遍历输出的键值是无序的，如果需要有序的键值对输出，需要对结果进行排序。
	for key, value := range m {
		fmt.Println(key, value)
	}

	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

	//匿名变量
	for _, value := range m {
		fmt.Println(value)
	}

}
