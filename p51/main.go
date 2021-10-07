package p51

import (
	"strings"
)

type Res struct {
	Val [][]int
}

func toString(r [][]int) [][]string {
	var res [][]string

	for _, v := range r {
		var sol []string
		for i := 0; i < len(v); i++ {
			var s strings.Builder
			for j := 0; j < len(v); j++ {
				if v[j] == i {
					s.WriteString("Q")
				} else {
					s.WriteString(".")
				}
			}
			sol = append(sol, s.String())
		}

		res = append(res, sol)
	}

	return res
}

func placeable(p []int, r, c, n int) bool {
	for i := 0; i < c; i++ {
		if p[i] == r {
			return false
		}

		if r >= i+1 && p[c-(i+1)] == r-(i+1) {
			return false
		}
		if r+i+1 < n && p[c-(i+1)] == r+i+1 {
			return false
		}
	}

	return true
}

func _solveNQueens_bt(res *Res, p []int, c int, n int) {
	if c == n {
		if len(p) == n {
			res.Val = append(res.Val, p)
		}
		return
	}

	for r := 0; r < n; r++ {
		if placeable(p, r, c, n) {
			pc := append([]int{}, p...)
			pc = append(pc, r)
			_solveNQueens_bt(res, pc, c+1, n)
		}
	}
}

func solveNQueens(n int) [][]string {
	res := &Res{}

	_solveNQueens_bt(res, nil, 0, n)

	return toString(res.Val)
}
