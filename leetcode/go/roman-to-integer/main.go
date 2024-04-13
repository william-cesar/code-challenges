func romanToInt(s string) int {
	var m = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	value := m[string(s[0])]

	for i := 1; i < len(s); i++ {
		prev := m[string(s[i-1])]
		curr := m[string(s[i])]

		if prev < curr {
			value += -(2 * prev) + curr
			continue
		}

		value += curr
	}

	return value
}