package main

import "fmt"

func main3() {
	var c1 complex64
	c1 = 1.10 + 0.1i

	// default complex128
	c2 := 1.10 + 0.1i

	c3 := complex(1.10, 0.1)

	c4 := complex(1, 1)

	fmt.Println(c1 == complex64(c2))

	c5 := complex128(c1)
	fmt.Println(c5 == c2) //false 由于c1是64位的，低位转高位，精度出现了问题，debug可发现

	c6 := complex128(c2)
	fmt.Println(c6 == c2) //true

	fmt.Println(c2 == c3)

	x := real(c4) //default float64
	y := imag(c4)

	fmt.Println(x)
	fmt.Println(y)
}
