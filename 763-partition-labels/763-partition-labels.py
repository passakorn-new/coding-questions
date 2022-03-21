class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        start, end = 0, 0
        results, last_seen = [], {}
        
        for i, char in enumerate(s):
            last_seen[char] = i

        for i, char in enumerate(s):
            end = max(end, last_seen[char])
            
            if i == end:
                results.append(end - start + 1)
                start = end + 1
            
        return results
