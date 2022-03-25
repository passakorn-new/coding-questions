class Solution:
    def twoCitySchedCost(self, costs: List[List[int]]) -> int:
        # sort from diff costs[i][0] with costs[i][1]
        sortedCosts = sorted(costs, key = lambda x: x[0] - x[1])
        res, N = 0, len(sortedCosts) // 2
        
        # send first two persons to A and rest to B
        for i in range(N):
            res += sortedCosts[i][0] + sortedCosts[i + N][1]
        return res
        