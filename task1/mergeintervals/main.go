package main

import (
	"fmt"
	"sort"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。


提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

func merge(intervals [][]int) [][]int {
	//根据左侧端点进行正序排序，
	//sort.Slice(intervals, func(i, j int) bool {
	//	return false
	//})
	//排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	//res = append(res, intervals[0])
	arr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= arr[1] {
			if intervals[i][1] > arr[1] { // 1,3   3,4/    1,5  2,3
				arr[1] = intervals[i][1]
			}
			continue
		}
		res = append(res, arr)
		arr = intervals[i]
	}
	res = append(res, arr)
	//遍历，根据右端点与下一个左端点进行比较，大于等于则合并，右端点前进1，
	return res
}

func main() {
	fmt.Println(merge([][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
