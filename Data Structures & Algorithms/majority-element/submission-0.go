func majorityElement(nums []int) int {
    counter := map[int]int{}
	n := len(nums) / 2
	for _, num := range nums {
		counter[num]++
		if counter[num] > n {
			return num
		}
	}
	return 0
}
