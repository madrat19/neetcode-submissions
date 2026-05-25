func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}
	for i, cur := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < cur {
			index := stack[len(stack)-1]
			ans[index] = i - index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
