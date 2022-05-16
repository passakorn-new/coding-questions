func shortestPathBinaryMatrix(grid [][]int) int {
    n := len(grid)
    
    if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
        return -1
    }
    
    directions := [8][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
    queue := [][3]int{{0, 0, 1}}
    
    for len(queue) > 0 {
        tuple := queue[0]
        queue = queue[1:]
        
        x, y, path := tuple[0], tuple[1], tuple[2]
        
        if x == n-1 && y == n-1 {
            return path
        }
        
        for _, dir := range directions {
            newX, newY := x + dir[0], y + dir[1]
            
            if newX >= 0 && newY >= 0 && newX < n && newY < n && grid[newX][newY] == 0 {
                queue = append(queue, [3]int{newX, newY, path + 1})
                grid[newX][newY] = 1
            }
        }
    }
    return -1
}