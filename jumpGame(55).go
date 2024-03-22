func canJump(nums []int) bool {
    max_val := 0

    for idx, val := range(nums) {
		if idx > max_val {
			return false
		}
		max_val = max(max_val, idx + val)
	}
	return true
}

func max(x, y int) int {
    if x >= y {
        return x
    }
    return y
}