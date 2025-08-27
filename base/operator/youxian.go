package main

import "fmt"

func TestYouXian() {
	var a int = 21
	var b int = 10
	var c int = 16
	var d int = 5
	var e int

	fmt.Printf("a,b,c,d,e = %d,%d,%d,%d,%d\n",
		a,
		b,
		c,
		d,
		e,
	)

	e = (a + b) * c / d // ( 31 * 16 ) / 5
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)

	e = ((a + b) * c) / d // ( 31 * 16 ) / 5
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n", e)

	e = (a + b) * (c / d) // 31 * (16/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)

	// 21 + (160/5)
	e = a + (b*c)/d
	fmt.Printf("a + (b * c) / d 的值为  : %d\n", e)

	// 2 & 2 = 2; 2 * 3 = 6; 6 << 1 = 12; 3 + 4 = 7; 7 ^ 3 = 4;4 | 12 = 12
	f := 3 + 4 ^ 3 | 2&2*3<<1
	fmt.Println(f == 12)
}
