package main

import "fmt"

// 十六进制
var a uint8 = 0xF
var b uint8 = 0xf

// 八进制
var c uint8 = 017
var d uint8 = 0o17
var e uint8 = 0o17

// 二进制
var f uint8 = 0b1111
var g uint8 = 0b1111

// 十进制
var h uint8 = 15

func main1() {
	fmt.Println(a, b, c, d, e, f, g, h)

	var float1 float32 = 10
	float2 := 10.0
	//float1 = float32(float2)
	fmt.Println(float1, float2)

	var aaa rune = '2'
	fmt.Println(string(aaa))
}
