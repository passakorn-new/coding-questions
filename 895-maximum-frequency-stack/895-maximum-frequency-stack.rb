class FreqStack
    def initialize()
        @hash_map = Hash.new { |h, k| h[k] = [] }
        @freq = Hash.new(0)
        @max_freq = 0
    end

    def push(val)
        @freq[val] += 1
        @max_freq = [@max_freq, @freq[val]].max
        @hash_map[@freq[val]] << val
    end

    def pop()
        remove_item = @hash_map[@max_freq].pop
        @freq[remove_item] -= 1
        @max_freq -= 1 if @hash_map[@max_freq].size == 0
        remove_item
    end
end

# Your FreqStack object will be instantiated and called as such:
# obj = FreqStack.new()
# obj.push(val)
# param_2 = obj.pop()