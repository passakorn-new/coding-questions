# @param {Integer} n
# @return {Integer[][]}

def generate_matrix(n)  
    c1, c2 = 0 , n - 1
    r1, r2 = 0, n - 1
    val = 1
    res = Array.new(n) { Array.new }
    
    while r1 <= r2 && c1 <= c2
        # moving left to right
        (c1..c2).step do |c|
            res[r1][c] = val
            val += 1
        end
        
        # moving down
        ((r1 + 1)..r2).step do |r|
            res[r][c2] = val
            val += 1
        end

        
        if r1 < r2 && c1 < c2
            # move right to left
            (c2 - 1).downto(c1) do |c|
                res[r2][c] = val
                val += 1
            end
            
            # move up
            (r2 - 1).downto(r1 + 1) do |r|
                res[r][c1] = val
                val += 1
            end
        end
        
        r1 += 1
        r2 -= 1
        c1 += 1
        c2 -= 1
    end
    
    res
end