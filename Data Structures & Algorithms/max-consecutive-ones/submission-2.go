func findMaxConsecutiveOnes(nums []int) int {
	maxS := 0
	curS := 0
	for _, num := range nums {
		if num == 1 {
			curS++
		} else {
			maxS = max(maxS, curS)
			curS = 0
		}
	}
	maxS = max(maxS, curS)
	return maxS
}
