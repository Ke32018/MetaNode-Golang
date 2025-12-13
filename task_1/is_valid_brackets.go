package task_1

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]rune, 0, len(s))
	pair := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, ch := range s {
		switch ch {
		case '{', '[', '(':
			stack = append(stack, ch)
		case '}', ']', ')':
			if len(stack) == 0 || stack[len(stack)-1] != pair[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
