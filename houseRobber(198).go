func rob(nums []int) int {
	maxNow := 0
	lastMax := 0
	for _, num := range nums {
		newMax := max(num+lastMax, maxNow)
		lastMax = maxNow
		maxNow = newMax
	}
	return maxNow
}