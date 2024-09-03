package algorithm

func LeetCode003(s string) int {
	var (
		result int
		start  int
	)

	m := make(map[byte]bool)

	for end := 0; end < len(s); end++ {
		for m[s[end]] {
			m[s[start]] = false
			start++
		}

		m[s[end]] = true

		if end-start+1 > result {
			result = end - start + 1
		}
	}

	return result
}
