package calculator

import "testing"

type AddData struct {
	x, y   int
	result int
}

// Single TestCase
func TestAdd(t *testing.T) {
	result := Add(5, 6)
	if result != 11 {
		t.Errorf("Add(5, 6) Failed. Expected %d, got %d\n", 11, result)
	} else {
		t.Logf("Add(5, 6) Success. Expected %d, got %d\n", 11, result)
	}
}

// Single TestCase
func TestSub(t *testing.T) {
	result := Sub(6, 5)

	if result != 1 {
		t.Errorf("Sub(6, 5) Failed. Expected %d, got %d\n", 1, result)
	} else {
		t.Logf("Sub(6, 5) Success. Expected %d, got %d\n", 1, result)
	}
}

func TestMultipleAdd(t *testing.T) {
	testData := []AddData{
		{1, 2, 3},
		{5, 7, 12},
		{5, -2, 3},
		{-2, -4, -6},
	}
	for _, data := range testData {
		result := Add(data.x, data.y)

		if result != data.result {
			t.Errorf("Add(%d, %d) Failed. Expected %d, got %d\n", data.x, data.y, data.result, result)
		} else {
			t.Logf("Add(%d, %d) Success. Expected %d, got %d\n", data.x, data.y, data.result, result)
		}
	}
}
