func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if fast == slow {
			break
		}
	}
	first, second := 0, slow
	for {
		first = nums[first]
		second = nums[second]
		if first == second {
			return first
		}
	}
}
