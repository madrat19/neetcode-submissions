func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := make([]int, 0)
	maxArea := 0

	for i, h := range heights {
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) != 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(height*width, maxArea)

		}
		stack = append(stack, i)
	}
	return maxArea
}
