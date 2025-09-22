package util

import "math"

func AbsI32(val int32) int32 {
	if val == math.MinInt32 {
		return math.MaxInt32
	}

	if val < 0 {
		return -val
	}

	return val
}
