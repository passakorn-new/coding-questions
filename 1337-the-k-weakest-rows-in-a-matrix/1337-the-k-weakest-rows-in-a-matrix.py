from queue import PriorityQueue

class Solution:    
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        q = PriorityQueue()
        
        for idx, m in enumerate(mat):
            # soldiers = sum(m)
            soldiers = self.binary_search(m) + 1
            q.put((soldiers, idx))

        res = []
        while k > 0:
            res.append(q.get()[1])
            k -= 1
        return res
    
    def binary_search(self, m: List[int]) -> int:
        l, r = 0, len(m) - 1
        
        while l <= r:
            mid = (l + r) // 2
            
            if (m[mid] != 0):
                l = mid + 1
            else:
                r = mid - 1
        return l
    
    
    # NOTES
    # - sum soldiers each row
    # - sort by soldier if same number of soldier consider index  (implement by priority queue)
    # - can improve when sum soldiers by use binary search for find 1 last index equal sum of soldiers