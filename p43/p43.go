package p43

func radd(n1, n2 []int) []int {
	var res []int
	i := 0
	p := 0
	for ; i < len(n1) && i < len(n2); i++ {
		t := n1[i] + n2[i] + p
		p = t / 10
		res = append(res, t%10)
	}

	n := n1
	if i < len(n2) {
		n = n2
	}

	for ; i < len(n); i++ {
		t := n[i] + p
		p = t / 10
		res = append(res, t%10)

		if p == 0 {
			res = append(res, n[i+1:]...)
			break
		}
	}

	if p != 0 {
		res = append(res, p)
	}

	return res
}

func rmulti(n []int, m int) []int {
	var res []int

	p := 0
	for i := range n {
		t := n[i]*m + p
		p = t / 10
		res = append(res, t%10)
	}

	if p != 0 {
		res = append(res, p)
	}

	return res
}

func multiply(num1 string, num2 string) string {
	var n1, n2 []int
	for i := len(num1) - 1; i >= 0; i-- {
		n1 = append(n1, int(num1[i]-'0'))
	}
	for i := len(num2) - 1; i >= 0; i-- {
		n2 = append(n2, int(num2[i]-'0'))
	}

	var val []int
	for i := range n1 {
		var p []int
		for j := 0; j < i; j++ {
			p = append(p, 0)
		}

		v := rmulti(n2, n1[i])
		v = append(p, v...)

		val = radd(val, v)
	}

	var res []rune
	first := true
	for i := len(val) - 1; i >= 0; i-- {

		// Remove '0' from the front end.
		if first && val[i] == 0 {
			continue
		}

		first = false
		res = append(res, rune(val[i]+'0'))
	}

	if len(res) == 0 {
		return "0"
	}

	return string(res)
}
