func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums) && j < len(nums); {
		if nums[i] == 0 && nums[j] != 0 {
			if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			} else {
				j++
			}
		} else {
			if nums[i] != 0 {
				i++
			}
			if nums[j] == 0 {
				j++
			}
		}
	}
}


