package algorithm

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() string {
	var str string
	for l != nil {
		str = fmt.Sprintf("%v%v", str, l.Val)
		if l.Next != nil {
			str = fmt.Sprintf("%v->", str)
		}

		l = l.Next
	}

	return str
}

func (l *ListNode) Create(vals []int) *ListNode {
	p := l
	for _, v := range vals {
		n := &ListNode{Val: v, Next: nil}
		p.Next = n
		p = p.Next
	}

	return l.Next
}
