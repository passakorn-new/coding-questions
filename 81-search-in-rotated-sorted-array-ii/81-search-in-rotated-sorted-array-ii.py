class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        left, right = 0, len(nums) - 1

        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                return True

            # case duplicate number
            if nums[left] == nums[mid] and nums[left] == nums[right]:
                left += 1
                right -= 1

            # first half is in order
            elif nums[left] <= nums[mid]:
                
                # target is in first  half
                if nums[left] <= target and nums[mid] > target:
                    right = mid - 1
                else:
                    left = mid + 1

            # second half is order
            else:
                
                # target is in second half
                if  nums[right] >= target and nums[mid] < target:
                    left = mid + 1
                else:
                    right = mid - 1
        return False  