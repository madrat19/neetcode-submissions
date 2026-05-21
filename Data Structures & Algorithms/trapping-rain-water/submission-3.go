func trap(height []int) int {
	l := 0
	r := len(height) - 1
	maxLeft, maxRight := height[0], height[r]
	ans := 0

	for l < r {
		if maxLeft < maxRight {
			l++
			ans += max(0, maxLeft-height[l])
			maxLeft = max(maxLeft, height[l])
		} else {
			r--
			ans += max(0, maxRight-height[r])
			maxRight = max(maxRight, height[r])
		}
	}
	return ans
}
