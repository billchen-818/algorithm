package algorithm

func LeetCode206(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head

	for cur != nil {
		// 	temp := cur.Next
		// 	cur.Next = pre
		// 	pre = cur
		// 	cur = temp
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}
