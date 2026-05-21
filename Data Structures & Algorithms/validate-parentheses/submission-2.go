func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	stack := []rune{}
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}
			if r == ')' {
				if stack[len(stack)-1] != '(' {
					return false
				}
				stack = stack[:len(stack)-1]
			}
			if r == ']' {
				if stack[len(stack)-1] != '[' {
					return false
				}
				stack = stack[:len(stack)-1]
			}
			if r == '}' {
				if stack[len(stack)-1] != '{' {
					return false
				}
				stack = stack[:len(stack)-1]
			}

		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}