package p445

/**
* Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}

	p := l
	c := l.Next
	p.Next = nil

	for c != nil {
		n := c.Next
		c.Next = p

		p = c
		c = n
	}

	return p
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r1 := reverse(l1)
	r2 := reverse(l2)

	head := &ListNode{}
	cur := head
	j := 0

	for r1 != nil && r2 != nil {
		v := r1.Val + r2.Val + j
		cur.Next = &ListNode{
			Val: v % 10,
		}
		j = v / 10

		cur = cur.Next
		r1 = r1.Next
		r2 = r2.Next

	}

	for r1 != nil {
		v := r1.Val + j
		cur.Next = &ListNode{
			Val: v % 10,
		}
		j = v / 10

		r1 = r1.Next
		cur = cur.Next
	}

	for r2 != nil {
		v := r2.Val + j
		cur.Next = &ListNode{
			Val: v % 10,
		}
		j = v / 10

		r2 = r2.Next
		cur = cur.Next
	}

	if j != 0 {
		cur.Next = &ListNode{
			Val: j,
		}
	}

	// delete head
	res := head.Next
	head.Next = nil

	return reverse(res)
}
