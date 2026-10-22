func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 1
	}
	k := 0
	for i := 0; i < l; i++ {
		if i+1 == l || nums[i+1] != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

