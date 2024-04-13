func tictactoe(moves [][]int) string {
	const TO_WIN = 3

	col := [2][3]int{{0, 0, 0}, {0, 0, 0}}
	row := [2][3]int{{0, 0, 0}, {0, 0, 0}}
	diag := [2][2]int{{0, 0}, {0, 0}}

	for i, v := range moves {
		var player = i % 2
		col[player][v[0]]++
		row[player][v[1]]++

		if v[0] == v[1] {
			diag[player][0]++
		}

		if v[0]+v[1] == 2 {
			diag[player][1]++
		}

		if col[player][0] == TO_WIN || col[player][1] == TO_WIN || col[player][2] == TO_WIN || row[player][0] == TO_WIN || row[player][1] == TO_WIN || row[player][2] == TO_WIN || diag[player][0] == TO_WIN || diag[player][1] == TO_WIN {
			if player == 0 {
				return "A"
			}
			return "B"
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}