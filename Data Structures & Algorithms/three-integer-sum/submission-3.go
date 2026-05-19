func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				k--
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return ans
}