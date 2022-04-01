class Solution:
    def reverseString(self, s: List[str]) -> None:
        left = 0
        right = len(s) - 1
        
        while left < right:
            temp = s[left]
            s[left] = s[right]
            s[right] = temp
            
            right -= 1
            left += 1
            
        # Two pointer