func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	var xor byte
	var sumT, sumS int

	for i, _ := range s {
		xor ^= s[i]
		xor ^= t[i]

		sumS += int(s[i] * s[i])
		sumT += int(t[i] * t[i])
	}

	return xor == 0 && sumT == sumS
}