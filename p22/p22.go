package p22

type res struct {
	r []string
}

func generateParenthesis(n int) []string {
	r := &res{}
	backtrack(r, "", 0, 0, n)
	return r.r
}

func backtrack(r *res, cur string, open, close, max int) {
	if len(cur) == max *2 {
		r.r = append(r.r, cur)
		return
	}

	if open < max {
		backtrack(r, cur+"(", open+1, close, max)
	}
	if close < open {
		backtrack(r, cur+")", open, close +1, max)
	}
}



