package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

*/
/*
思路：遍历字符串，出现左括号就压入栈中，出现右括号就判断栈顶是否匹配，匹配则出栈，不匹配则返回false
*/
func isValid(s string) bool {

	// 创建一个空的int切片
	stack := make([]int, 0, 100)
	//var stack []int  // nil切片
	//stack := []int{} // 空切片（非nil）
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, '(')
			// fallthrough 是 Go 语言中的一个关键字，用于在 switch 语句中显式地继续执行下一个 case 分支。
			// 通常情况下，Go 中的 switch 语句会在每个 case 执行完毕后自动跳出，不会继续执行下面的 case。
			// 使用 fallthrough 可以让程序继续执行紧随其后的下一个 case 或 default 分支，而不需要重新匹配条件。
			// 注意：fallthrough 必须是某个 case 最后一条语句，并且不能用在 default 分支中。
		case ')':
			//当出现右括号的时候，栈里的最后一个元素必须是对应的左括号
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			//栈里最后一个元素出栈
			stack = stack[:len(stack)-1]
		case '{':
			stack = append(stack, '{')
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '[':
			stack = append(stack, '[')
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			panic("invalid char")
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false

}

func main() {
	//定义一个字符串切片
	//s := "12345"
	//
	//a := []rune(s)
	//a1 := a[1:]
	//fmt.Println(string(a1))
	//字符串截取，是大于等于左边下标，小于右边下标
	//a2 := a[:len(a)-1]
	//fmt.Println(a2)
	//fmt.Println(string(a2))

	fmt.Println(isValid("{[]}"))

}
