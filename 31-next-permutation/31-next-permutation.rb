# @param {Integer[]} nums
# @return {Void} Do not return anything, modify nums in-place instead.

def next_permutation(nums)
    
    # find index before sort desc
    pointer = nums.length - 2
    while pointer >= 0 && nums[pointer] >= nums[pointer + 1]
        pointer -= 1
    end
    
    # case nums is sorted asc
    return nums.reverse! if pointer == -1
    
    # find index for swap with pointer
    greater = nums.length - 1
    while greater > pointer && nums[greater] <= nums[pointer]
        greater -= 1
    end
    
    # swap
    nums[greater], nums[pointer] = nums[pointer], nums[greater]
    
    # reverse
    nums[pointer + 1..] = nums[pointer + 1..].reverse!
end

    # NOTES ex [0 ,1 ,2, 10, 9, 4, 3, 0]
    # 1. desc = [10, 9, 4, 3, 0]. so pointer is 2
    # 2. greater = last index of nums = 7
        # 2.1 find index nums[pointer] < nums[greater] and subtract greater 1 each loop
        # 2.2    (2 < 0) greater 7 false -> (2 < 3) greater 6 true
    # 3. swap index pointer with index greater
    # 4. reverse nums since index pointer + 1
    
    # reference https://www.youtube.com/watch?v=4wlBBRo4tYY