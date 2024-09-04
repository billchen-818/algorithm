package algorithm

func LeetCode202(num int) bool {
	m := make(map[int]bool)

	for {
		nextValue := 0

		for num > 0 {
			nextValue += (num % 10) * (num % 10)
			num /= 10
		}

		if _, ok := m[nextValue]; ok {
			return nextValue == 1
		}

		m[nextValue] = true
		num = nextValue
	}
}
