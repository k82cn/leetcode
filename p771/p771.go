package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func numJewelsInStones(J string, S string) int {
	num := 0
	for _, s := range S {
		for _, j := range J {
			if j == s {
				num++
			}
		}
	}
	return num
}
