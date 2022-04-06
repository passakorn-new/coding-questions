MOD = 10**9 + 7

# @param {Integer[]} arr
# @param {Integer} target
# @return {Integer}
def three_sum_multi(arr, target)
    res, freq_nums = 0, Array.new(101, 0)
    arr.each { |num| freq_nums[num] += 1 }
    
    (0..100).step do |i|
        (i..100).step do |j|
            # k is remain value to equal target
            k = target - i - j
            
            # next if k is impossible value  (0 <= arr[i] <= 100)
            next if k < 0 || k > 100
            
            if i == j && j == k
                # pp "A" 
                # pp "i = #{i}, j = #{j}, k = #{k}"
                # pp "res = #{(freq_nums[i] * (freq_nums[i]-1) * (freq_nums[i]-2)) /6}"
                res += (freq_nums[i] * (freq_nums[i]-1) * (freq_nums[i]-2)) /6;
            elsif i == j && j != k
                # pp "B"
                # pp "i = #{i}, j = #{j}, k = #{k}"
                # pp "res = #{((freq_nums[i] * (freq_nums[i]-1)) / 2) * freq_nums[k]}"
                res += ((freq_nums[i] * (freq_nums[i]-1)) / 2) * freq_nums[k];
            elsif i < j && j < k
                # pp "C" 
                # pp "i = #{i}, j = #{j}, k = #{k}"
                # pp "res = #{ freq_nums[i] * freq_nums[j] * freq_nums[k]}"
                res += freq_nums[i] * freq_nums[j] * freq_nums[k]
            end
            
            # pp "-------------------"
            res %= 1000000007
        end
    end
    
    res
end