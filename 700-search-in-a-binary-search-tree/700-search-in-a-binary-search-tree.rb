# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val = 0, left = nil, right = nil)
#         @val = val
#         @left = left
#         @right = right
#     end
# end
# @param {TreeNode} root
# @param {Integer} val
# @return {TreeNode}

def search_bst(root, val)
    while root && root.val != val do
        if root.val < val
            root = root.right
        else
            root = root.left
        end
    end
    
    root
end