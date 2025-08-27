package main

import "fmt"

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。

示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
加 1 后得到 123 + 1 = 124。
因此，结果应该是 [1,2,4]。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
加 1 后得到 4321 + 1 = 4322。
因此，结果应该是 [4,3,2,2]。
示例 3：

输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 9。
加 1 得到了 9 + 1 = 10。
因此，结果应该是 [1,0]。

提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits 不包含任何前导 0。
*/
func plusOne1(digits []int) []int {
	//res := make([]int, 1, 100)
	//从右想左遍历,最右侧如果是9则替换为0，非9则加1返回
	jinyi := false
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			jinyi = false
			return digits
		}
		if digits[i] == 9 {
			jinyi = true
			digits[i] = 0
		}
	}
	if jinyi {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func plusOne(digits []int) []int {
	//res := make([]int, 1, 100)
	//从右想左遍历,最右侧如果是9则替换为0，非9则加1返回
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		if digits[i] == 9 {
			digits[i] = 0
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{1, 2, 0}))
	fmt.Println(plusOne([]int{9, 8, 9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{9}))
}
