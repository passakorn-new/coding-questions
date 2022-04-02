class Solution:
    def validPalindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1
        
        while left < right:
            if s[left] != s[right]:
                return self.isPalindome(s, left + 1, right) or self.isPalindome(s, left, right - 1)
            else:
                left += 1
                right -= 1
        return True
    
    def isPalindome(self, s, left, right) -> bool:
        while left < right:
            if s[left] != s[right]:
                return False
            else:
                left += 1
                right -= 1
        return True