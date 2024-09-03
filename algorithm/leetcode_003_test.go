package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode003(t *testing.T) {
	assert.Equal(t, LeetCode003("abcabcbb"), 3)
	assert.Equal(t, LeetCode003("bbbbb"), 1)
	assert.Equal(t, LeetCode003("pwwkew"), 3)

}
