package main

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

func IsBalanced(root *TreeNode) bool {
	_, balanced := balancedHeight(root)
	return balanced
}

func balancedHeight(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	lh, lb := balancedHeight(root.Left)
	rh, rb := balancedHeight(root.Right)
	if !lb || !rb {
		return -1, false
	}
	if lh >= rh {
		if lh-rh > 1 {
			return -1, false
		}
		return 1 + lh, true
	} else if rh-lh > 1 {
		return -1, false
	}
	return 1 + rh, true
}
