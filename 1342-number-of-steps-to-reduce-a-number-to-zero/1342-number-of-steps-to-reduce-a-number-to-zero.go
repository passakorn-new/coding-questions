
func numberOfSteps(num int) int {
    if num <= 1 {
        return num
    }
    
    var c int
    for num > 0 {
        c += num & 1 + 1
        num >>= 1
    }
    
    return c - 1
}