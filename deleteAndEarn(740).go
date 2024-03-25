func deleteAndEarn(nums []int) int {
	h := map[int]int{}
	for _, n := range nums {
		if _, ok := h[n]; ok {
			h[n] += 1
		} else {
			h[n] = 1
		}
	}
	arr := make([]int, len(h))
	dp := make([]int, len(h))
	i := 0
	for key, _ := range h {
		arr[i] = key
		i += 1
	}
	slices.Sort(arr)
	dp[0] = arr[0] * h[arr[0]]
	for i := 1; i < len(arr); i++ {
		n := arr[i]
		dp[i] = n * h[n]
		if n-1 == arr[i-1] {
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
		} else {
			dp[i] += dp[i-1]
		}
		dp[i] = max(dp[i], dp[i-1])
	}
	return dp[len(dp)-1]
}