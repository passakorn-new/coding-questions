MOD = 10**9 + 7

# @param {Integer[]} arr
# @param {Integer} target
# @return {Integer}
def three_sum_multi(arr, target)
  result = 0
  arr.sort!

  arr.count.times do |i|
    t = target - arr[i]
    j = i + 1
    k = arr.count - 1

    while j < k
      if arr[j] + arr[k] < t
        j += 1
      elsif arr[j] + arr[k] > t
        k -= 1
      elsif arr[j] != arr[k]
        left = right = 1
        
        while j + 1 < k && arr[j] == arr[j+1]
          left += 1
          j += 1
        end
        
        while k - 1 > j && arr[k] == arr[k-1]
          right += 1
          k -= 1
        end

        result += left*right
        result %= MOD
        j += 1
        k -= 1
      else
        result += (k - j + 1)*(k - j)/2
        result %= MOD
        break
      end
    end
  end
  
  result
end