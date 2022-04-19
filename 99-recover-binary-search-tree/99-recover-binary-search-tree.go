/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
We need to find the first and second elements that are not in order right?

How do we find these two elements? For example, we have the following tree that is printed as in order traversal:

6, 3, 4, 5, 2

We compare each node with its next one and we can find out that 6 is the first element to swap because 6 > 3 and 2 is the second element to swap because 2 < 5.

Really, what we are comparing is the current node and its previous node in the "in order traversal".

cc https://leetcode.com/problems/recover-binary-search-tree/discuss/32535/No-Fancy-Algorithm-just-Simple-and-Powerful-In-Order-Traversal
*/

func recoverTree(root *TreeNode)  {
    var prevEle, firstEle, secondEle *TreeNode
    INT_MIN := -1 * int(1e10)
    prevEle = &TreeNode{ Val: INT_MIN, Left: nil, Right: nil }
    
    // In order traversal to find the two elements
    traverse(root, &prevEle, &firstEle, &secondEle)
    
    // Swap the values of the two nodes
    if firstEle != nil && secondEle!= nil{
        firstEle.Val, secondEle.Val = secondEle.Val, firstEle.Val
    }
}

func traverse (root *TreeNode, prevEle, firstEle, secondEle **TreeNode)  {
    if root == nil {
        return
    }
    
    traverse(root.Left, prevEle, firstEle, secondEle)
    
    // If first element has not been found, assign it to prevElement (refer to 6 in the example above)
    if *firstEle == nil && (*prevEle).Val >= root.Val{
        *firstEle = *prevEle    
    }
    
    // If first element is found, assign the second element to the root (refer to 2 in the example above)
    if *firstEle != nil && (*prevEle).Val >= root.Val{
        *secondEle = root    
    }
    
    *prevEle = root
    traverse(root.Right, prevEle, firstEle, secondEle)
}