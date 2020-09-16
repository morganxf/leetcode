package tree_and_graph

import "testing"

func Test_numIslans_3(t *testing.T) {
	var intput [][]byte = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	out := numIslans_3(intput)
	expected := 1
	if out != expected {
		t.Errorf("expected: %v, got: %v", expected, out)
	}
}
