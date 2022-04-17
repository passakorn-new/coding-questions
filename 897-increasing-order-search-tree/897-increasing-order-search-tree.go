/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// result = inorder(root.left) + root + inorder(root.right)
func increasingBST(root *TreeNode) *TreeNode {
    return dfs(root, nil)
}


func dfs(root, tail *TreeNode) *TreeNode {
    // tail = previous node
    if root == nil {
        return tail
    }
    
    res := dfs(root.Left, root)
	root.Left = nil
	root.Right = dfs(root.Right, tail)
    return res
}