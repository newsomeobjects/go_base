package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为int: %d \n", num)
	str1 := strconv.Itoa(num)
	fmt.Printf("int转换为字符串: %s \n", str1)

	ui64, err := strconv.ParseUint(str, 10, 32)
	fmt.Printf("字符串转换为uint64: %d \n", num)

	str2 := strconv.FormatUint(ui64, 2)
	fmt.Printf("uint64转换为字符串: %s \n", str2)

	v := "42"
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseUint(v, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	//The bitSize argument specifies the integer type that the result must fit into.
	//Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64.
	//If bitSize is below 0 or above 64, an error is returned.
	h := "-100"
	//自动类型转换 - 即使指定32位，函数仍然返回int64类型，但会检查是否超出32位范围
	if s, err := strconv.ParseInt(h, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} else {
		panic(err)
	}

	h1 := "2147483648" // 超出32位正数范围
	// 会返回错误 32主要是为了数值范围验证，确保字符串表示的数值在32位整数的有效范围内。
	if s, err := strconv.ParseInt(h1, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} else {
		//panic(err)
	}

	var i interface{} = 3
	a, ok := i.(int)
	if ok {
		fmt.Printf("'%d' is a int \n", a)
	} else {
		fmt.Println("conversion failed")
	}
	switch v := i.(type) {
	case int:
		fmt.Println("i is a int", v)
	case string:
		fmt.Println("i is a string", v)
	default:
		fmt.Println("i is unknown type", v)
	}
}
