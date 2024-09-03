package algorithm

func LeetCode001(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
		if i, ok := m[target-value]; ok {
			return []int{i, index}
		}
		m[value] = index
	}

	return []int{}
}
