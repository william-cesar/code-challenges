func judgeCircle(moves string) bool {
	dir := map[rune]int{
		'L': 0,
		'R': 0,
		'U': 0,
		'D': 0,
	}

	for _, v := range moves {
		dir[v]++
	}

	return dir['L'] == dir['R'] && dir['U'] == dir['D']
}