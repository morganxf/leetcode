package array

import "testing"

func TestLongestPalindrome_1(t *testing.T) {
	input := "cbbd"
	out := LongestPalindrome_1(input)
	if out != "bb" {
		t.Errorf("expected: bb. got: %s", out)
	}
}
