func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	l := len(arr)
	diff := arr[l-1] - arr[l-2]

	for i := l - 1; i > 0; i-- {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true
}