package array

import (
	"fmt"
	"testing"
)

func TestSum3Nums_1(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	out := Sum3Nums_1(input)
	fmt.Println(out)
}

func TestSum3Nums_2(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	out := Sum3Nums_2(input)
	fmt.Println(out)
}
