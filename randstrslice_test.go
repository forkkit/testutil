package testutil_test

import (
	"testing"

	"github.com/openacid/testutil"
	"github.com/stretchr/testify/require"
)

func TestRandStrSlice(t *testing.T) {

	ta := require.New(t)

	strs := testutil.RandStrSlice(1000, 5, 10)
	for _, s := range strs {

		ta.GreaterOrEqual(len(s), 5)
		ta.Less(len(s), 10)

		for _, c := range []byte(s) {
			ta.True(isIn(testutil.Runes, c))
		}
	}
}

func TestRandStrSlice_specifyRunes(t *testing.T) {

	ta := require.New(t)

	strs := testutil.RandStrSlice(10, 5, 10, []rune("12"))
	for _, s := range strs {

		ta.GreaterOrEqual(len(s), 5)
		ta.Less(len(s), 10)

		for _, c := range []byte(s) {
			ta.True(c == byte('1') || c == byte('2'))
		}
	}
}

func TestRandBytesSlice(t *testing.T) {

	ta := require.New(t)

	bss := testutil.RandBytesSlice(1000, 5, 10)
	for _, s := range bss {

		ta.GreaterOrEqual(len(s), 5)
		ta.Less(len(s), 10)

		for _, c := range s {
			ta.True(isIn(testutil.Runes, c))
		}
	}
}

func TestRandBytesSlice_specifyRunes(t *testing.T) {

	ta := require.New(t)

	bss := testutil.RandBytesSlice(10, 5, 10, []rune("12"))
	for _, s := range bss {

		ta.GreaterOrEqual(len(s), 5)
		ta.Less(len(s), 10)

		for _, c := range s {
			ta.True(c == '1' || c == '2')
		}
	}
}

func isIn(collection []rune, c byte) bool {
	for _, r := range collection {
		if c == byte(r) {
			return true
		}
	}

	return false
}
