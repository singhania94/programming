package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(root *TreeNode) {
	if root == nil {
		return
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invert(root.Left)
	invert(root.Right)
}
