package p8

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	data := []rune(strings.TrimSpace(s))
	// If empty string, return 0
	if len(data) == 0 {
		return 0
	}

	positive := true

	switch data[0] {
	case '-', '+':
		if data[0] == '-' {
			positive = false
		}
		data = data[1:]
	}

	// If empty string, return 0
	if len(data) == 0 {
		return 0
	}

	if !unicode.IsDigit(data[0]) {
		return 0
	}

	max := int64(math.MaxInt32)
	if !positive {
		max = -math.MinInt32
	}

	var res int64
	for i := range data {
		if !unicode.IsDigit(data[i]) {
			break
		}

		res *= 10
		res += int64(data[i] - '0')

		if res >= max {
			res = max
			break
		}
	}

	if !positive {
		return -int(res)
	}

	return int(res)
}
