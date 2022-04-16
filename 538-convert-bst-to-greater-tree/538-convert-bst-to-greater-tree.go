/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    dfs(root, &sum)
    return root
}

func dfs(node *TreeNode, sum *int) {
    if node == nil {
        return
    }

    // move to the far right
    dfs(node.Right, sum)

    // use pointer keep value to sum
    *sum += node.Val
    
    // change value in node
    node.Val = *sum
    
    // move to left
    dfs(node.Left, sum)
}