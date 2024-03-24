func maxArea(height []int) int {
	maxVolume, left, right := 0, 0, len(height)-1
	for left < right {
		volume := min(height[left], height[right]) * (right - left)
		if volume > maxVolume {
			maxVolume = volume
		}
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}
	return maxVolume
}