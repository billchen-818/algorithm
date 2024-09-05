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

func LeetCode002(l1 *ListNode, l2 *ListNode) *ListNode {
	var temp int

	l := &ListNode{}
	p := l

	for l1 != nil && l2 != nil {
		t := &ListNode{}
		t.Val = (temp + l1.Val + l2.Val) % 10
		t.Next = nil

		p.Next = t
		p = p.Next

		temp = (temp + l1.Val + l2.Val) / 10

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		t := &ListNode{}
		t.Val = (temp + l1.Val) % 10
		t.Next = nil

		p.Next = t
		p = p.Next

		temp = (temp + l1.Val) / 10

		l1 = l1.Next
	}

	for l2 != nil {
		t := &ListNode{}
		t.Val = (temp + l2.Val) % 10
		t.Next = nil

		p.Next = t
		p = p.Next

		temp = (temp + l2.Val) / 10

		l2 = l2.Next
	}

	if temp != 0 {
		t := &ListNode{}
		t.Val = temp
		t.Next = nil

		p.Next = t
	}

	return l.Next
}
