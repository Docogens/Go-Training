package main

import (
	"testing"
)

func Test_Solution(t *testing.T) {
	inputs := []struct {
		arr    []int
		result int64
	}{
		// adding more cases later
		{arr: []int{1, 2, 3, 4}, result: 0},
		{arr: []int{25, 35, 872, 228, 53, 278, 872}, result: 4},
	}

	for _, item := range inputs {
		result := solution(item.arr)

		if result != item.result {
			t.Errorf("result is %q, expected %q", item.result, result)
		}
	}

	/*a := []int{1, 2, 3, 4}
	var expected int64 = 0
	result := solution(a)

	if result != expected {
		t.Errorf("result is %q, expected %q", result, expected)
	}*/
}
