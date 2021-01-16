package testutil

import (
	"math/rand"
	"sort"
)

// Runes is a collection of all visibel ascii chars.
var Runes = []rune("~!@#$%^&*()_+`-=[]{};:<>?,./abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// Letters is a collection of all 0-9a-zA-Z
var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// RandStrSlice creates a sorted []string of var-length random string.
//
// The random string length is in range `[minLen, maxLen)`.
//
// One could specify the set from which to choose a random character with the
// 4th arg. It should be a []rune. Be default it choose character from
// Runes.
//
// NOTE:
// Assumes the length of `runes` is n,
// The total number of all possible random string is:
//
//   maxLen
//   ∑ n^i
//   i=minLen
//
// Thus it is possible this func blocks for ever with a very small source.
// E.g. the following statement could generate at most 5 different string:
//
//   RandStrSlice(10, 5, 10, []rune("1"))
//
// Since 0.1.3
func RandStrSlice(n, minLen, maxLen int, options ...interface{}) []string {

	var from []rune

	if len(options) > 0 {
		from = options[0].([]rune)
	}

	if from == nil {
		from = Runes
	}

	rlen := len(from)

	mp := make(map[string]bool)

	for i := 0; i < n; i++ {
		l := rand.Intn(maxLen-minLen) + minLen
		b := make([]rune, l)
		for j := 0; j < l; j++ {
			k := rand.Intn(rlen)

			b[j] = from[k]
		}
		s := string(b)
		if _, ok := mp[s]; ok {
			i--
		} else {
			mp[s] = true
		}
	}

	rst := make([]string, 0, n)
	for k := range mp {
		rst = append(rst, k)
	}

	sort.Strings(rst)
	return rst
}

// RandBytesSlice is similar to RandStrSlice, but it returns a slice of
// []byte instead of a slice of string.
//
// Since 0.1.3
func RandBytesSlice(n, minLen, maxLen int, options ...interface{}) [][]byte {
	strs := RandStrSlice(n, minLen, maxLen, options...)

	rst := make([][]byte, 0, len(strs))
	for _, s := range strs {
		rst = append(rst, []byte(s))
	}
	return rst
}
