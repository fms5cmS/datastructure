package stack

// 20.符号匹配：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 使用切片模拟栈
	stack := make([]rune, 0)
	top := -1
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			// 入栈，并移动栈顶指针
			stack = append(stack, v)
			top++
		case ')', ']', '}':
			if top < 0 {
				return false
			}
			// 获取栈顶元素，并进行匹配
			if m[v] != stack[top] {
				return false
			}
			// 匹配成功，移除栈顶元素
			stack = stack[:top]
			top--
		}
	}
	// 判断栈是否为空，如果不为空，则符号没有完全匹配
	return top == -1
}
