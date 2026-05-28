import "slices"

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, slices.Max(piles)
	for l < r {
		m := (l + r) / 2
		if time(piles, m) > h {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func time(piles []int, speed int) int {
	time := 0
	for _, p := range piles {
		time += (p + speed - 1) / speed
	}
	return time
}
