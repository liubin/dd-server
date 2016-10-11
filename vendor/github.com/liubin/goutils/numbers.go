package goutils

import (
	"fmt"
	"strconv"
)

func Float64ToInt64(f float64) int64 {
	s := fmt.Sprintf("%.f", f)
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
