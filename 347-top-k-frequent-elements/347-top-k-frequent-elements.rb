# @param {Integer[]} nums
# @param {Integer} k
# @return {Integer[]}

def top_k_frequent(nums, k)
    # bucket sort

    res = []
    freq_nums = Hash.new(0)
    bucket = Array.new(nums.length + 1)
    
    # accumulate freq number
    nums.each { |num| freq_nums[num] += 1 }
    
    # store num in bucket by freq
    freq_nums.each { |num, freq| bucket[freq].nil? ? bucket[freq] = [num] : bucket[freq] << num }
    
    # search k freq from last index in bucket
    bucket.length.downto(1) do |i|
        res += bucket[i] if bucket[i] != nil
        
        return res if res.length == k
    end
end