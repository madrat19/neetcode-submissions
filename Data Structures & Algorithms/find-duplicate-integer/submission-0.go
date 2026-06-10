func findDuplicate(nums []int) int {
	for _, n := range nums {
		if nums[abs(n)] < 0 {
			return abs(n)
		}
		nums[abs(n)] *= -1
	}
	return 0
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
