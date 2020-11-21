package testutil_test

import (
	"testing"

	"github.com/openacid/testutil"
	"github.com/stretchr/testify/require"
)

func TestRandI32Slice(t *testing.T) {

	ta := require.New(t)

	i32s := testutil.RandI32Slice(10, 100, 3)

	ta.Equal(100, len(i32s))
	ta.GreaterOrEqual(i32s[0], int32(10))

	for i, i32 := range i32s {
		if i == 0 {
			continue
		}

		ta.GreaterOrEqual(i32, i32s[i-1])
		ta.GreaterOrEqual(i32-i32s[i-1], int32(0))
		ta.LessOrEqual(i32-i32s[i-1], int32(3))

	}

}

func TestRandU32Slice(t *testing.T) {

	ta := require.New(t)

	u32s := testutil.RandU32Slice(10, 100, 3)

	ta.Equal(100, len(u32s))
	ta.GreaterOrEqual(u32s[0], uint32(10))

	for i, u32 := range u32s {
		if i == 0 {
			continue
		}

		ta.GreaterOrEqual(u32, u32s[i-1])
		ta.GreaterOrEqual(u32-u32s[i-1], uint32(0))
		ta.LessOrEqual(u32-u32s[i-1], uint32(3))

	}

}

func TestRandI64Slice(t *testing.T) {

	ta := require.New(t)

	i64s := testutil.RandI64Slice(10, 100, 3)

	ta.Equal(100, len(i64s))
	ta.GreaterOrEqual(i64s[0], int64(10))

	for i, i64 := range i64s {
		if i == 0 {
			continue
		}

		ta.GreaterOrEqual(i64, i64s[i-1])
		ta.GreaterOrEqual(i64-i64s[i-1], int64(0))
		ta.LessOrEqual(i64-i64s[i-1], int64(3))

	}

}

func TestRandU64Slice(t *testing.T) {

	ta := require.New(t)

	u64s := testutil.RandU64Slice(10, 100, 3)

	ta.Equal(100, len(u64s))
	ta.GreaterOrEqual(u64s[0], uint64(10))

	for i, u64 := range u64s {
		if i == 0 {
			continue
		}

		ta.GreaterOrEqual(u64, u64s[i-1])
		ta.GreaterOrEqual(u64-u64s[i-1], uint64(0))
		ta.LessOrEqual(u64-u64s[i-1], uint64(3))

	}

}
