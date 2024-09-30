package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(nums)
	expected := 5
	if result != expected {
		t.Errorf("removeDuplicates failed; expected %d, got %d", expected, result)
	}
}
