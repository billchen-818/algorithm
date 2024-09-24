package algorithm

import (
	"reflect"
	"sort"
)

// 问题描述
// 三数之和

func Leetcode015(nums []int) [][]int {
	//res := solution1(nums)
	res := solution2(nums)
	return res
}

func _solution1(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	newRes := deleteRepetionItem(res)

	return newRes
}

func deleteRepetionItem(items [][]int) [][]int {

	newItems := make([][]int, 0)
	for i := 0; i < len(items); i++ {
		flag := false

		for j := i + 1; j < len(items); j++ {
			if equal(items[i], items[j]) {
				flag = true
				break
			}
		}

		if !flag {
			newItems = append(newItems, items[i])
		}

	}

	return newItems
}

func equal(a, b []int) bool {

	return reflect.DeepEqual(a, b)
}

// func equal(a []int, b []int) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for i := 0; i < len(a); i++ {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 		continue
// 	}

// 	return true
// }

func solution2(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if v, ok := m[-nums[i]-nums[j]]; ok {
				if v == 1 && (nums[i] != nums[j] && nums[j] != -nums[i]-nums[j]) {
					res = append(res, genint(nums[i], nums[j], -nums[j]-nums[i]))
				}

				if v == 2 && (nums[i] == nums[j] || nums[i] == -nums[i]-nums[j] || nums[j] == -nums[j]-nums[i]) {
					res = append(res, genint(nums[i], nums[j], -nums[j]-nums[i]))
				}

				if v == 3 && (nums[i] == nums[j] && nums[i] == 0) {
					res = append(res, genint(nums[i], nums[j], -nums[j]-nums[i]))
				}
			}
		}
	}

	newRes := deleteRepetionItem(res)

	return newRes

}

func genint(a, b, c int) []int {
	if a >= b {
		if b >= c {
			return []int{c, b, a}
		} else if a >= c {
			return []int{b, c, a}
		} else {
			return []int{b, a, c}
		}
	} else {
		if a >= c {
			return []int{c, a, b}
		} else if b >= c {
			return []int{a, c, b}
		} else {
			return []int{a, b, c}
		}
	}
}
