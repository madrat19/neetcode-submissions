func removeElement(nums []int, val int) int {
	swapped := 0
	j := len(nums) - 1
	for i := 0; i <= j; i++ {
		if nums[j] == val {
			for j > -1 && nums[j] == val {
				j--
				swapped ++
			}
		}
		if nums[i] == val && i < j {
			nums[i], nums[j] = nums[j], nums[i]
			swapped ++
			j --
		}
	}
	return len(nums) - swapped
}
