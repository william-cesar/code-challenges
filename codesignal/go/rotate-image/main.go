/*
[00, 01, 02]    [20, 10, 00]
[10, 11, 12] -> [21, 11, 01]
[20, 21, 22]    [22, 12, 02]


[00, 01, 02, 03]    [30, 20, 10, 00]
[10, 11, 12, 13] -> [31, 21, 11, 01]
[20, 21, 22, 23]    [32, 22, 12, 02]
[30, 31, 32, 33]    [33, 23, 13, 03]
*/

func solution(a [][]int) [][]int {

	rotated := make([][]int, 0)

	for idx := range a {
		start := len(a) - 1
		row := make([]int, 0)

		for start >= 0 {
			row = append(row, a[start][idx])
			start--
		}
		rotated = append(rotated, row)
	}

	return rotated
}