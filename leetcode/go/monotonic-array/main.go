func isMonotonic(nums []int) bool {
	dir := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			if dir == 0 {
				dir = -1
			} else if dir == 1 {
				return false
			}
		} else if nums[i] > nums[i+1] {
			if dir == 0 {
				dir = 1
			} else if dir == -1 {
				return false
			}
		}
	}

	return true
}