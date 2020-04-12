package p17

var num = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

type res struct {
	r []string
}

func letterCombinations(digits string) []string {
	r := &res{}

	backtrack(r, "", digits, 0)

	return r.r
}

func backtrack(r *res, cur, org string, i int) {
	if len(org) == 0 {
		return
	}

	if i == len(org) {
		r.r = append(r.r, cur)
		return
	}

	c := org[i]
	ss := num[string(c)]
	for _, s := range ss {
		backtrack(r, cur+s, org, i+1)
	}
}
