package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		tmp := p
		p = q
		q = tmp
	}

	if p.Val <= root.Val && q.Val >= root.Val {
		return root
	}
	if p.Val > root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}
