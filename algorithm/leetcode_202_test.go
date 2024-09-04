package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode202(t *testing.T) {
	assert.Equal(t, true, LeetCode202(19))
	assert.Equal(t, false, LeetCode202(2))
}
