func strStr(haystack string, needle string) int {
	hLen := len(haystack)
	nLen := len(needle)
	start := -1
	limit := hLen - nLen

	for i := 0; i < limit+1; i++ {
		if haystack[i:i+nLen] == needle {
			start = i
			break
		}
	}

	return start
}