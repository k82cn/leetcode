package p24

import "testing"

func buildList(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	head := &ListNode{Val: data[0]}
	prev := head
	for i := 1; i < len(data); i++ {
		prev.Next = &ListNode{Val: data[i]}
		prev = prev.Next
	}
	return head
}

func TestSwapPairs(t *testing.T) {
	cases := []struct {
		arg1     *ListNode
		expected *ListNode
	}{
		{
			buildList([]int{1, 2, 3, 4}),
			buildList([]int{2, 1, 4, 3}),
		},
	}

	for i, c := range cases {
		got := swapPairs(c.arg1)
		if !equals(got, c.expected) {
			t.Errorf("case %d failed: expected %v, got %v", i, got, c.expected)
		}
	}
}
