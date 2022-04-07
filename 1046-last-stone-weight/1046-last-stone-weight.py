from sortedcontainers import SortedList

class Solution(object):
    def lastStoneWeight(self, stones):
        stones_sorted = SortedList(stones)
        while len(stones_sorted) >= 2:
            y = stones_sorted.pop()
            x = stones_sorted.pop()
            
            if y > x: stones_sorted.add(y - x) # stones_sorted is sorted list
                
        return stones_sorted.pop() if len(stones_sorted) else 0