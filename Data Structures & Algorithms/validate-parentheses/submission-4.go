func isValid(s string) bool {
	stack := []rune{}
	closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, r := range s {
		if open, ok := closeToOpen[r]; ok {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}