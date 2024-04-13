func arraySign(nums []int) int {
	res := 1

	for _, v := range nums {
		if v == 0 {
			return v
		}

		if v < 0 {
			res *= -1
		}
	}

	return res
}