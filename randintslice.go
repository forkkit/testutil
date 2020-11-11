package testutil

import (
	"math/rand"
	"time"
)

// RandI32Slice generates a sorted int32 slice of `n` elts in it, starting from
// `min`, and every two adjacent elts differ from 0 to `step`.
func RandI32Slice(min, n, step int32) []int32 {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	rst := make([]int32, 0)

	p := min
	for i := 0; i < int(n); i++ {
		s := int32(rnd.Float64() * float64(step))
		p += s
		rst = append(rst, p)
	}

	return rst
}
