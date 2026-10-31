func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := 0; i < len(flowerbed); i++ {
		if (i - 1 == -1 || flowerbed[i-1] == 0) &&
		 (i+1 == len(flowerbed) || flowerbed[i+1] == 0) &&
		 flowerbed[i] == 0 {
			flowerbed[i] = 1
			n--
		 }
	}
	return n <= 0
}