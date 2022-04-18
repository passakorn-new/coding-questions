/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    increseOrderSearchTree := dfs(root, nil)
    
    for k - 1 > 0 {
		increseOrderSearchTree = increseOrderSearchTree.Right
        k--
	}
    
    return increseOrderSearchTree.Val
}

func dfs(root *TreeNode, tail *TreeNode) *TreeNode {
    if root == nil {
        return tail
    }
    
    // recursive call, traversing left while passing in the current node as tail
    res := dfs(root.Left, root)
    
    // we don't want the current node to have a left child, only a single right child
	root.Left = nil
    
    // tail is parent's parent from line 25
	root.Right = dfs(root.Right, tail)

    return res
}