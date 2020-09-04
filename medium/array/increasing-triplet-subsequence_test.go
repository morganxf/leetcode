package array

import "testing"

func TestIncreasingTriplet_1(t *testing.T) {
	input := []int{5, 1, 5, 5, 2, 5, 4}
	var expected bool = true
	out := IncreasingTriplet_1(input)
	if out != expected {
		t.Errorf("expected: %t, got: %t", expected, out)
	}
}
