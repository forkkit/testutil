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

// RandU32Slice generates a sorted int32 slice of `n` elts in it, starting from
// `min`, and every two adjacent elts differ from 0 to `step`.
func RandU32Slice(min uint32, n, step int32) []uint32 {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	rst := make([]uint32, 0)

	p := min
	for i := 0; i < int(n); i++ {
		s := uint32(rnd.Float64() * float64(step))
		p += s
		rst = append(rst, p)
	}

	return rst
}

// RandI64Slice generates a sorted int64 slice of `n` elts in it, starting from
// `min`, and every two adjacent elts differ from 0 to `step`.
func RandI64Slice(min, n, step int64) []int64 {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	rst := make([]int64, 0)

	p := min
	for i := 0; i < int(n); i++ {
		s := int64(rnd.Float64() * float64(step))
		p += s
		rst = append(rst, p)
	}

	return rst
}

// RandU64Slice generates a sorted uint64 slice of `n` elts in it, starting from
// `min`, and every two adjacent elts differ from 0 to `step`.
func RandU64Slice(min, n, step uint64) []uint64 {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	rst := make([]uint64, 0)

	p := min
	for i := 0; i < int(n); i++ {
		s := uint64(rnd.Float64() * float64(step))
		p += s
		rst = append(rst, p)
	}

	return rst
}
