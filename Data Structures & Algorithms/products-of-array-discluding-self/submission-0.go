func productExceptSelf(nums []int) []int {
	l := len(nums)
	leftPrefix := make([]int, l)
	rightPrefix := make([]int, l)
	leftPrefix[0] = 1
	rightPrefix[l-1] = 1
	product := 1
	for i := 0; i < l-1; i++ {
		product *= nums[i]
		leftPrefix[i+1] = product
	}
	product = 1
	for i := l - 1; i > 0; i-- {
		product *= nums[i]
		rightPrefix[i-1] = product
	}
	ans := make([]int, l)
	for i := 0; i < l; i++ {
		ans[i] = leftPrefix[i] * rightPrefix[i]
	}
	return ans
}
