package tree_and_graph

// 如果存在重复的值会怎么样？
func buildTree(preorder []int, inorder []int) *TreeNode {
	//return buildTree_1(preorder, inorder)
	//return buildTree_2(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
	return buildTree_3(preorder, inorder)
}

func buildTree_1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	var i int = 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func buildTree_2(preorder []int, inorder []int, preorderLeft, preorderRight, inorderLeft, inorderRight int) *TreeNode {
	if preorderLeft > preorderRight {
		return nil
	}
	rootVal := preorder[preorderLeft]
	root := &TreeNode{Val: rootVal}
	var i int = inorderLeft
	for ; i <= inorderRight; i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	root.Left = buildTree_2(preorder, inorder, preorderLeft+1, i-inorderLeft+preorderLeft, inorderLeft, i-1)
	root.Right = buildTree_2(preorder, inorder, i-inorderLeft+preorderLeft+1, preorderRight, i+1, inorderRight)
	return root
}

// 迭代
func buildTree_3(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var stack []*TreeNode
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		curVal := preorder[i]
		preNode := stack[len(stack)-1]
		if preNode.Val != inorder[inorderIndex] {
			// 树中的左节点
			preNode.Left = &TreeNode{Val: curVal}
			stack = append(stack, preNode.Left)
		} else {
			// 如果 stackCur.Val != inorder[inorderIndex], 则说明 stackPre存在右孩子
			// if len(stack) == 0, 则说明stackPre是root节点
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				preNode = stack[len(stack)-1]
				// preNode从stack中移除
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			preNode.Right = &TreeNode{Val: curVal}
			stack = append(stack, preNode.Right)
		}
	}
	return root
}
