func carFleet(target int, position []int, speed []int) int {
	posSpeed := make(map[int]int, len(speed))
	for i, _ := range position {
		posSpeed[position[i]] = speed[i]
	}
	sort.Ints(position)

	fleets := 1
	head := position[len(position)-1]
	for i := len(position) - 2; i > -1; i-- {
		if float32(target-position[i])/float32(posSpeed[position[i]]) > float32((target-head))/float32(posSpeed[head]) {
			head = position[i]
			fleets++
		}
	}
	return fleets
}
