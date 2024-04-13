func findTheDifference(s string, t string) byte {
	var s1, s2 int

	for _, v := range s {
		s1 += int(v)
	}

	for _, v := range t {
		s2 += int(v)
	}

	d := s2 - s1
	return byte(d)
}