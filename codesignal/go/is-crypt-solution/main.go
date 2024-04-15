func solution(crypt []string, solution [][]string) bool {
	kvMap := make(map[string]string)

	for _, v := range solution {
		kvMap[v[0]] = v[1]
	}

	v1, l1 := processMessage(crypt[0], kvMap)
	v2, l2 := processMessage(crypt[1], kvMap)
	v3, l3 := processMessage(crypt[2], kvMap)

	if l1 || l2 || l3 {
		return false
	}

	return v1+v2 == v3
}

func processMessage(msg string, mp map[string]string) (int64, bool) {
	val := ""

	for _, v := range msg {
		val = val + mp[string(v)]
	}

	sum, _ := strconv.ParseInt(val, 10, 64)

	hasLeadingZero := len(msg) > 1 && mp[string(msg[0])] == "0"

	return sum, hasLeadingZero
}
