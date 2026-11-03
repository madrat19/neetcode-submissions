func minSubArrayLen(target int, nums []int) int {
	sum := nums[0]
	ans := len(nums) + 1
	for l, r := 0, 0; r < len(nums) && l < len(nums); {
		if sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		} else {
			r++
			if r < len(nums) {
				sum += nums[r]
			}
		}
	}
	if ans == len(nums)+1 {
		ans = 0
	}
	return ans
}
