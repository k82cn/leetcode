package p53

import "testing"

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		arg      []int
		expected int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
	}

	for i, c := range cases {
		got := maxSubArray(c.arg)
		if got != c.expected {
			t.Errorf("case %d failed: expected %d, got %d", i, c.expected, got)
		}
	}
}
