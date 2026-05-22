func evalRPN(tokens []string) int {
	var a, b int
	stack := []int{}
	for _, token := range tokens {
		n, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, n)
		} else {
			a = stack[len(stack)-2]
			b = stack[len(stack)-1]
			switch token {
			case "+":
				a = a + b
			case "-":
				a = a - b
			case "*":
				a = a * b
			case "/":
				a = a / b
			}
			stack[len(stack)-2] = a
			stack = stack[:len(stack)-1]
		}

	}
	return stack[0]
}