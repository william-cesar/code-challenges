func calPoints(operations []string) int {
	var record []int

	for _, v := range operations {
		l := len(record)

		if v == "D" {
			record = append(record, record[l-1]*2)
		} else if v == "C" {
			record = record[:l-1]
		} else if v == "+" {
			record = append(record, record[l-2]+record[l-1])
		} else {
			i, _ := strconv.Atoi(v)
			record = append(record, i)
		}
	}

	res := 0

	for _, v := range record {
		res = res + v
	}

	return res
}