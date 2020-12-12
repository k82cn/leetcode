package p43

import "testing"

func TestMultiply(t *testing.T) {
	cases := []struct {
		arg1, arg2 string
		expected   string
	}{
		{
			"2", "3", "6",
		},
		{
			"123", "456", "56088",
		},
		{
			"9133", "0", "0",
		},
	}

	for i, c := range cases {
		got := multiply(c.arg1, c.arg2)
		if c.expected != got {
			t.Errorf("case %d failed: exptected %s, got %s", i, c.expected, got)
		}
	}
}
