const maxNumber = 9
const toNumber = 49

func isFilled(grid [][]string, block [maxNumber]bool, start, end int) bool {
	return block[grid[start][end][0]-toNumber]
}

func solution(grid [][]string) bool {
	const subGridSize = 3

	for i := 0; i < maxNumber; i++ {
		var row, col, subGrid [maxNumber]bool

		for j := 0; j < maxNumber; j++ {
			if grid[i][j] != "." {
				if isFilled(grid, row, i, j) {
					return false
				}

				row[grid[i][j][0]-toNumber] = true
			}

			if grid[j][i] != "." {
				if isFilled(grid, col, j, i) {
					return false
				}
				col[grid[j][i][0]-toNumber] = true
			}

			subRow := i%subGridSize*subGridSize + j%subGridSize
			subCol := i/subGridSize*subGridSize + j/subGridSize

			if grid[subCol][subRow] != "." {

				if isFilled(grid, subGrid, subCol, subRow) {
					return false
				}

				subGrid[grid[subCol][subRow][0]-toNumber] = true
			}
		}
	}
	return true
}