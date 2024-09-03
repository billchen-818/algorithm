package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode001(t *testing.T) {
	assert.Equal(t, LeetCode001([]int{2, 7, 11, 15}, 9), []int{0, 1})
	assert.Equal(t, LeetCode001([]int{4, 3, 9}, 12), []int{1, 2})
	assert.Equal(t, LeetCode001([]int{3, 4, 6, 15}, 9), []int{0, 2})
}
