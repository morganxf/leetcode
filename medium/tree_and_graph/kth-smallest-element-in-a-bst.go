package tree_and_graph

func kthSmallest(root *TreeNode, k int) int {
	//return kthSmallest_1(root, k)
	return kthSmallest_2(root, k)
}

func kthSmallest_1(root *TreeNode, k int) int {
	var arr []int
	inorder(root, &arr)
	if len(arr) < k {
		return -1
	}
	return arr[k-1]
}

func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

func kthSmallest_2(root *TreeNode, k int) int {
	var stack []*TreeNode
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return curNode.Val
		}
		root = curNode.Right
	}
}
