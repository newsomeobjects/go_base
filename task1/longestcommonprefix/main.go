package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
提示：
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 如果非空，则仅由小写英文字母组成
*/

/*
思路：遍历字符串数组，找到所有字符串的最短长度，然后遍历最短长度的字符串，判断所有字符串是否相同。
*/
func longestCommonPrefix1(stars []string) string {
	var res []byte
	if len(stars) == 1 {
		return stars[0]
	}
	for i := 0; i < len(stars); i++ {
		s := stars[i]
		if len(s) == 0 {
			return ""
		}

		if len(res) == 0 && i < len(stars)-1 { //第一个字符串,放到res中
			//res = append(res, s[0:len(res)])
			res = append(res, s[0:]...)
			continue
		}
		if len(res) > len(s) { //公共前缀长度比当前字符串长，截取
			res = res[:len(s)]
			//res = append(res, s[:len(res)]...)
		}

		for j := 0; j < len(res); j++ { //遍历最短长度的字符串
			if res[j] != s[j] {
				res = res[:j]
				if len(res) == 0 {
					return ""
				}
				break
			}
		}

	}

	return string(res)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	res := []byte{}
loop1:
	for { //按列遍历所有字符的第i个元素
		var c byte //当前遍历的元素
		for _, s := range strs {
			//取字符串的第i个字符
			if i >= len(s) {
				break loop1
			}
			//v := []byte(s)
			d := []byte(s)[i]
			if c == 0 {
				c = d
			} else if c != d {
				break loop1
			}
		}
		res = append(res, c)
		i++
	}
	return string(res)
}

/*
首先检查数组是否为空，为空则返回空字符串
以第一个字符串为基准，逐个字符比较
外层循环遍历第一个字符串的每个字符位置
内层循环检查其他字符串在相同位置是否匹配
如果发现不匹配或到达某个字符串末尾，则返回当前已匹配的前缀部分
如果所有字符都匹配，则返回第一个字符串作为公共前缀
*/
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] { //后续字符串长度不够，或者字符不匹配
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(longestCommonPrefix([]string{"flower"}))

	//fmt.Println(longestCommonPrefix1([]string{"ca", "c", "bba", "bacb", "bcb"}))
	//var res []byte
	//var c byte
	fmt.Println(longestCommonPrefix([]string{"bba", "bacb", "bcb"}))
}
