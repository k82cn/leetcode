package main

import "fmt"

func main() {
	fmt.Printf("+%v\n", sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	res := []int{}
	for _, a := range A {
		if a&1 == 0 {
			res = append([]int{a}, res...)
		} else {
			res = append(res, a)
		}
	}
	return res
}
