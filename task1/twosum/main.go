package main

import "fmt"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]


提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
*/

/*
思路：爆破穷举，n(n-1)/2,j=i+1
*/

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		//if nums[i] > target {
		//	continue
		//}
		e := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == e {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		k, exist := m[nums[i]]
		if exist {
			return []int{k, i}
		}
		e := target - nums[i]
		m[e] = i //key放差值，value放当前坐标
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 2, 4}, 8))
	fmt.Println(twoSum([]int{-1, -2, -3, -4, -5}, -8))
	//nums := []int{3, 4, 5}
	//m := make(map[int]int)
	//fmt.Println(m)
	//fmt.Println(m[1000])

}
