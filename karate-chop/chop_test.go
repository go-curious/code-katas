package karate_chop_test

import (
	"testing"

	karate_chop "github.com/go-curious/code-katas/karate-chop"
)

type TestTable []struct {
	i   int
	arr []int
	exp int
}

func TestChop(t *testing.T) {
	table := TestTable{
		{i: 3, arr: []int{}, exp: -1},
		{i: 3, arr: []int{1}, exp: -1},
		{i: 1, arr: []int{1}, exp: 0},

		{i: 1, arr: []int{1, 3, 5}, exp: 0},
		{i: 3, arr: []int{1, 3, 5}, exp: 1},
		{i: 5, arr: []int{1, 3, 5}, exp: 2},
		{i: 0, arr: []int{1, 3, 5}, exp: -1},
		{i: 2, arr: []int{1, 3, 5}, exp: -1},
		{i: 4, arr: []int{1, 3, 5}, exp: -1},
		{i: 6, arr: []int{1, 3, 5}, exp: -1},

		{i: 1, arr: []int{1, 3, 5, 7}, exp: 0},
		{i: 3, arr: []int{1, 3, 5, 7}, exp: 1},
		{i: 5, arr: []int{1, 3, 5, 7}, exp: 2},
		{i: 7, arr: []int{1, 3, 5, 7}, exp: 3},
		{i: 0, arr: []int{1, 3, 5, 7}, exp: -1},
		{i: 2, arr: []int{1, 3, 5, 7}, exp: -1},
		{i: 4, arr: []int{1, 3, 5, 7}, exp: -1},
		{i: 6, arr: []int{1, 3, 5, 7}, exp: -1},
		{i: 8, arr: []int{1, 3, 5, 7}, exp: -1},
	}

	for _, test := range table {
		result := karate_chop.Chop(test.i, test.arr, 0, len(test.arr)-1)
		if test.exp != result {
			t.Errorf("Test failed. Expected: %d, Received: %d", test.exp, result)
		}
	}
}
