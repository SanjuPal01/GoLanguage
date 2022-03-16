package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositiveSort(t *testing.T) {
	tests := []struct {
		Name     string
		Numbers  []int
		Expected []int
	}{
		{
			Name:     "Test1",
			Numbers:  []int{5, 3, 6, 3, 6, 1},
			Expected: []int{1, 3, 3, 5, 6, 6},
		},
		{
			Name:     "Test2",
			Numbers:  []int{4, 4, 4, 4, 4, 4},
			Expected: []int{4, 4, 4, 4, 4, 4},
		},
		{
			Name:     "Test3",
			Numbers:  []int{-1, -2, -6, -3, -4},
			Expected: []int{-6, -4, -3, -2, -1},
		},
		{
			Name:     "Test4",
			Numbers:  []int{0, 2, 3, -2, 5, 3},
			Expected: []int{-2, 0, 2, 3, 3, 5},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			t.Logf(test.Name)
			res := Sort(test.Numbers)
			// assert.True(t, assert.ObjectsAreEqualValues(test.Expected, res))
			assert.Equal(t, test.Expected, res)
		})
	}
}

func TestNegativeSort(t *testing.T) {
	tests := []struct {
		Name     string
		Numbers  []int
		Expected []int
	}{
		{
			Name:     "Test1",
			Numbers:  []int{5, 3, 6, 3, 6, 1},
			Expected: []int{1, 3, 3, 6, 5, 6},
		},
		{
			Name:     "Test2",
			Numbers:  []int{4, 4, 4, 4, 4, 4},
			Expected: []int{4, 4, 4, 3, 4, 4},
		},
		{
			Name:     "Test3",
			Numbers:  []int{-1, -2, -6, -3, -4},
			Expected: []int{-6, -4, -2, -3, -1},
		},
		{
			Name:     "Test4",
			Numbers:  []int{0, 2, 3, -2, 5, 3},
			Expected: []int{-2, 1, 2, 3, 3, 5},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			t.Logf(test.Name)
			res := Sort(test.Numbers)
			// assert.True(t, assert.ObjectsAreEqualValues(test.Expected, res))
			assert.NotEqual(t, test.Expected, res)
		})
	}
}
