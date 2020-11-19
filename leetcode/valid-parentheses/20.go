/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := []rune{}

    for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ')')
		} else if ch == '[' {
			stack = append(stack, ']')
		} else if ch == '{' {
			stack = append(stack, '}')
		} else {
			if stackLength := len(stack); stackLength < 1 || ch != stack[stackLength - 1] {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}

	return len(stack) == 0
}
// @lc code=end

