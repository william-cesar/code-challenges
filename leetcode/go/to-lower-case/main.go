func toLowerCase(s string) string {
	output := make([]byte, 0, len(s))

	for _, ch := range []byte(s) {

		if ch >= 'A' && ch <= 'Z' {
			ch += ' '
		}

		output = append(output, ch)
	}

	return string(output)
}