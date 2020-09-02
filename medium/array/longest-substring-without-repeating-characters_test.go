package array

import (
	"testing"
)

func TestLengthOfLongestSubstring_2(t *testing.T) {
	input := "tmmzuxt"
	out := LengthOfLongestSubstring_2(input)
	if out != 5 {
		t.Errorf("expected %d, but got: %d", 5, out)
	}
}
