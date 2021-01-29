package testutil

import (
	"math/bits"

	"github.com/openacid/low/mathext/zipf"
)

// Workload creates a workload that accesses an array with `n` items.
// The returned array contains which data to access at every step, according to
// workload type.
// The length is at least `n` and is power of two if not specified.
//
// Supported workload type is "zipf" and "scan".
//
// Since 0.1.4
func Workload(typ string, n int, opts ...int) []int {
	times := leastPow2Mask(n) + 1
	if len(opts) > 0 {
		times = opts[0]
	}

	if typ == "zipf" {
		return zipf.Accesses(1, 1.5, n, times, nil)
	} else if typ == "scan" {
		return accessesScan(n, times)
	}

	panic("unknown workload type:" + typ)
}

func accessesScan(n int, times int) []int {
	a := make([]int, times)
	for i := 0; i < times; i++ {
		a[i] = i % n
	}
	return a
}

func leastPow2Mask(n int) int {
	x := bits.LeadingZeros(uint(n))
	mask := uint(0)
	mask = (mask - 1) >> uint(x)
	return int(mask)
}

func maxMask(n int) int {
	mask := 0
	for ; (mask<<1 | 1) <= n; mask = mask<<1 | 1 {
	}
	return mask
}
