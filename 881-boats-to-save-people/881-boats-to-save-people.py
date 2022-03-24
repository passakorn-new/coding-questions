class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        # minimum boat = each boat should carry almost equal limit
        # implement left, right for use two pointer for group people
        count_boat, left, right = 0, 0, len(people) - 1 
        people.sort() # can improve if use count sort
        
        while left <= right:            
            if people[left] + people[right] <= limit:
                left += 1
                right -= 1
            else:
                right -= 1

            count_boat += 1
            
        return count_boat