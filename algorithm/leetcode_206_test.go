package algorithm

import (
	"fmt"
	"testing"
)

func TestLeetcode206(t *testing.T) {
	l := new(ListNode)
	l = l.Create([]int{1, 2, 3, 4, 5})
	fmt.Println(l.Print())

	l = LeetCode206(l)
	fmt.Println(l.Print())
}
