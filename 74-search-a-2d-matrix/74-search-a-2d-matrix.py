class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        n = len(matrix)
        m = len(matrix[0])
        left, right = 0, n * m - 1
        
        while left <= right:
            mid = (left + right) // 2
            mid_value = matrix[mid // m][mid % m]
            
            if mid_value == target:
                return True
            elif mid_value > target:
                right = mid - 1
            else:
                left = mid + 1
        return False