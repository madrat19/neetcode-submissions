func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	q := []int{}
	l, r := 0, 0
	for ; r < k; r++ {
		for len(q) > 0 && nums[r] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
	}
	ans = append(ans, nums[q[0]])
	for ; r < len(nums); l, r = l+1, r+1 {
		if q[0] == l {
			q = q[1:]
		}
		for len(q) > 0 && nums[r] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		ans = append(ans, nums[q[0]])

	}
	return ans
}
