func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
	for i, n := range nums {
		if j, ok := hash[target - n]; ok {
			return []int{j, i}
		}
		hash[n] = i
	}
	return []int{}
}
