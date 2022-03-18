# @param {String} s
# @return {Integer}
def score_of_parentheses(s)
    res = depth = 0
    stack = []
    
    s.each_char.with_index do |c, i|
        if c == '('
            depth += 1
        else
            depth -= 1
        end
        
        if s[i] == ')' && s[i - 1] == '('
            res += 2 ** depth
        end
    end
    
    res
end