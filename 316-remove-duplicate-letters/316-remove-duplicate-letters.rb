# @param {String} s
# @return {String}
def remove_duplicate_letters(s)
    freq = Array.new(26, 0)
    visited = Array.new(26, false)
    
    stack = []
    s.each_byte do |ch|
        freq[ch - 97] += 1
    end
    
    s.each_char do |ch|
        index = ch.ord - 97
        freq[index] -= 1
        next if visited[index]
        
        while !stack.empty? && stack[-1] > ch && freq[stack[-1].ord - 97 ] > 0 do
            visited[stack.pop.ord - 97] = false
        end
        
        visited[index] = true
        stack << ch
    end
    
    stack.join
end