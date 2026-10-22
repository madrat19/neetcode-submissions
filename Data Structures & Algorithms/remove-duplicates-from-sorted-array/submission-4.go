func removeDuplicates(nums []int) int {
	l := len(nums)
	k := 0
	for i := 0; i < l; i++ {
		if i+1 == l || nums[i+1] != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

