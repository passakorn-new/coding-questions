class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        left, right = 0, len(nums) - 1
        
        while left <= right:
            mid = (left + right) // 2
            cnt = 0
            
            for num in nums:
                if num <= mid:
                    cnt += 1
                    
            if cnt <= mid:
                left = mid + 1
            else:
                right = mid - 1
        
        return left
            
#         freq = {}
#         for num in nums:
#             if freq.get(num) != None:
#                 return  num
#             else:
#                 freq[num] =  1
        