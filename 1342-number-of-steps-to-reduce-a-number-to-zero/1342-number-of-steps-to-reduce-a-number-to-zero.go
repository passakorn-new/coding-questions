import (
    "math"
)

func numberOfSteps(num int) int {
    count := 0
    for num > 0 {
        if math.Mod(float64(num), 2) != 0 {
            num = num - 1
            count ++
        } else {
            num = num / 2
            count ++
        }
    }
    
    return count
}