package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode002(t *testing.T) {
	l1 := &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}

	l2 := &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}

	assert.Equal(t, "7->0->8", LeetCode002(l1, l2).Print())
}
