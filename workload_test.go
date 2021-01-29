package testutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLeastPow2Mask(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input int
		want  string
	}{
		{0, "0"},
		{1, "1"},
		{2, "11"},
		{3, "11"},
		{4, "111"},
	}

	for i, c := range cases {
		got := leastPow2Mask(c.input)
		gotStr := fmt.Sprintf("%b", got)
		ta.Equal(c.want, gotStr, "%d-th: case: %+v", i+1, c)
	}
}

func TestMaxMask(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input int
		want  string
	}{
		{0, "0"},
		{1, "1"},
		{2, "1"},
		{3, "11"},
		{4, "11"},
		{5, "11"},
		{6, "11"},
		{7, "111"},
	}

	for i, c := range cases {
		got := maxMask(c.input)
		gotStr := fmt.Sprintf("%b", got)
		ta.Equal(c.want, gotStr, "%d-th: case: %+v", i+1, c)
	}
}
