/*
"abacabad" -> {97:4 98:2 99:1 100:1}
find first with value 1 and return, else return "_"
*/

func solution(s string) string {
	kvMap := map[rune]int{}

	for _, r := range s {
		kvMap[r]++
	}

	for _, r := range s {
		if kvMap[r] == 1 {
			return string(r)
		}
	}

	return "_"
}