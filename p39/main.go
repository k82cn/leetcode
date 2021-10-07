package p39

import "sort"

type Val struct {
	res [][]int
}

func _combinationSum_bt(v *Val, sum int, cur []int, candidates []int, target int) {
	if sum == target {
		v.res = append(v.res, cur)
		return
	}

	for i := range candidates {
		if sum+candidates[i] > target {
			break
		}

		var c []int
		c = append(c, cur...)
		c = append(c, candidates[i])
		_combinationSum_bt(v, sum+candidates[i], c, candidates[i:], target)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	v := &Val{}
	sort.Ints(candidates)

	_combinationSum_bt(v, 0, nil, candidates, target)

	return v.res
}
