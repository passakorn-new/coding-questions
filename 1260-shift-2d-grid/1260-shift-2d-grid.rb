# @param {Integer[][]} grid
# @param {Integer} k
# @return {Integer[][]}
def shift_grid(grid, k)
    rows, col = grid.length, grid[0].length
    dimension = rows * col
    res = Array.new(rows) { Array.new }
    
    k %= dimension
    start = dimension - k
    current_index = 0
    
    (start...(start + dimension)).step do |i|
        n = (i / col) % rows
        m = i % col
        res[current_index / col] << grid[n][m]
        
        current_index += 1
    end
    
    res
end
