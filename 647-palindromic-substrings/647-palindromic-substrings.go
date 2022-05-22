func countSubstrings(s string) int {
    result := 0
    
    for i := range s {
        result += countPalindrome(s, i, i) + countPalindrome(s, i, i+1)
    }
    
    return result
}

func countPalindrome(s string, left, right int) int {
    count := 0
    
    for left >= 0 && right < len(s) {
        if s[left] != s[right] {
            return count
        }
        
        count++
        left--
        right++
    }
    
    return count
}