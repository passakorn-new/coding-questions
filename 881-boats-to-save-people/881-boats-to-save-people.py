class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        # minimum boat = each boat should carry almost equal limit
        # implement left, right for use two pointer for group people
        count_boat, left, right = 0, 0, len(people) - 1 
        # people.sort() # can improve if use count sort
        
        index, count = 0, [0] * (limit + 1)
        for p in people:
            count[p] += 1
            
        for val in range(1, limit + 1):
            while count[val] > 0:
                people[index] = val
                index += 1
                count[val] -= 1
        
        while left <= right:            
            if people[left] + people[right] <= limit:
                left += 1
                right -= 1
            else:
                right -= 1

            count_boat += 1
            
        return count_boat
    