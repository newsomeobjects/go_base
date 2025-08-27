package main

import "fmt"

/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
示例 1 ：
输入：nums = [2,2,1]
输出：1
示例 2 ：
输入：nums = [4,1,2,1,2]
输出：4
示例 3 ：
输入：nums = [1]
输出：1
*/

/*
思路：將切片中的数字放到map中，如果map中存在，則將map中該數字刪除，否則將該數字添加到map中，最终map中剩一個數字，就是只出現一次的數字
*/
func singleNumber1(nums []int) int {

	m := make(map[int]int)
	for _, v := range nums {
		//fmt.Println(i, v)
		//如果存在则删掉，不存在则添加
		if _, exists := m[v]; exists {
			delete(m, v)
			continue
		}
		m[v] = 1
	}

	//返回map中剩下的唯一一个元素
	for k := range m {
		return k
	}
	//map中没有元素说明出错了
	//return -1
	panic("no element")
}

/*
思路：异或运算，相同数字异或结果为0，0与数字异或结果为数字本身
*/
func singleNumber(nums []int) int {
	//遍历n-1个值，每个值和下一个值做异或运算
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
		//res ^= nums[i]
	}
	return res
}

func main() {
	fmt.Println(singleNumber1([]int{2, 2, 1, 3, 3, 5, 5, 6, 6}))
	fmt.Println(singleNumber([]int{2, 2, 1, 3, 3, 5, 5, 6, 6}))
}
