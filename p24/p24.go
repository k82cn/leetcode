package p24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	virt := &ListNode{Next: head}
	prev := virt
	cur := head
	next := head.Next

	for {
		prev.Next = next
		cur.Next = next.Next
		next.Next = cur

		if cur.Next == nil || cur.Next.Next == nil {
			break
		}

		prev = cur
		cur = prev.Next
		next = cur.Next
	}

	return virt.Next
}
