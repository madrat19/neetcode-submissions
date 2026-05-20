func maxArea(heights []int) int {
	ans := 0
	i, j := 0, len(heights)-1
	for i < j {
		area := (j - i) * min(heights[i], heights[j])
		ans = max(ans, area)
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}
