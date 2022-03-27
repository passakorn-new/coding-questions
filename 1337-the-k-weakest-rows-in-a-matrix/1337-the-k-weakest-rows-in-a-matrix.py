from queue import PriorityQueue

class Solution:
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        q = PriorityQueue()
        
        for idx, m in enumerate(mat):
            soldiers = sum(m)
            q.put((soldiers, idx))

        res = []
        while k > 0:
            res.append(q.get()[1])
            k -= 1
        return res
    
    
    # NOTES
    # - sum soldiers each row
    # - sort by soldier if same number of soldier consider index  (implement by priority queue)