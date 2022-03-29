class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        freq = {}
        for num in nums:
            if freq.get(num) != None:
                return  num
            else:
                freq[num] =  1
        