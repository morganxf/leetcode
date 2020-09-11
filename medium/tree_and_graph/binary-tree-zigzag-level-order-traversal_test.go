package tree_and_graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_zigzagLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	out := zigzagLevelOrder(root)
	expected := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}
	assert.Equal(t, expected, out)
}
