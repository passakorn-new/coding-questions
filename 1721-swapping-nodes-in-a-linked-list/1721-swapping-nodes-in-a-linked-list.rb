# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val = 0, _next = nil)
#         @val = val
#         @next = _next
#     end
# end
# @param {ListNode} head
# @param {Integer} k
# @return {ListNode}
def swap_nodes(head, k)
    left = right = curr = head
    
    count = 1
    while curr
        if count < k
            left = left.next
        end
        
        if count > k
            right = right.next
        end
        
        curr = curr.next
        count += 1
    end
    
    left.val, right.val = right.val, left.val
    head
end

    # Notes
    # Assume [1,2,3,4,5,6,7],  k=2
    # left start 0, right start 0
