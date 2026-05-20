func trap(height []int) int {
	l := len(height)
	maxLefts := make([]int, l)
	maxRights := make([]int, l)

	maxLeft := 0
	for i, h := range height {
		maxLefts[i] = maxLeft
		maxLeft = max(maxLeft, h)
	}

	maxRight := 0
	for i := l - 1; i > -1; i-- {
		maxRights[i] = maxRight
		maxRight = max(maxRight, height[i])
	}

	ans := 0
	for i := 0; i < l; i++ {
		water := min(maxLefts[i], maxRights[i]) - height[i]
		if water > 0 {
			ans += water
		}
	}
	return ans
}