/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BFS
func deepestLeavesSum(root *TreeNode) int {
    sumPreviousLeaveLevel := 0
    queue := []*TreeNode{}
    queue = append(queue, root)
    
    for len(queue) > 0 {
        l := len(queue)
        sumPreviousLeaveLevel = 0

        // loop queue same level
        for i := 0; i < l; i++ {
            currentNode := queue[0]
            sumPreviousLeaveLevel += currentNode.Val

            queue = queue[1:]
            
            if currentNode.Left != nil {
                queue = append(queue, currentNode.Left)
            }
            
            if currentNode.Right != nil {
                queue = append(queue, currentNode.Right)
            }
        }
    }

    return sumPreviousLeaveLevel
}
