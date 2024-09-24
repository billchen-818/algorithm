package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode015(t *testing.T) {
	t1 := []int{-1, 0, 1, 2, -1, -4}
	t2 := []int{0, 1, 1}
	t3 := []int{0, 0, 0}

	r1 := [][]int{{-1, 0, 1}, {-1, -1, 2}}
	r2 := [][]int{}
	r3 := [][]int{{0, 0, 0}}

	a1 := Leetcode015(t1)
	a2 := Leetcode015(t2)
	a3 := Leetcode015(t3)

	assert.Equal(t, r1, a1)
	assert.Equal(t, r2, a2)
	assert.Equal(t, r3, a3)
}
