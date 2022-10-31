func shipWithinDays(weights []int, days int) int {
    max, sum := maxAndSumWeight(weights)
    least_capacity := binarySearch(max, sum, days, weights)

    return least_capacity
}

func maxAndSumWeight(weights []int) (int, int){
    max, sum := 0, 0
    for _, weight := range weights {
        if max < weight {
            max = weight
        }
        sum += weight
    }
    
    return max, sum
}

func binarySearch (left int, right int, days int, weights []int) int {
    for left < right {
        mid := (left + right) / 2
        used_days := calConveyDays(weights, mid)
        
        if used_days > days {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return left
}

func calConveyDays(weights []int, capacity int) int{
    countDay, sum := 0, 0
    for _, weight := range weights {
        sum += weight
        if sum > capacity {
            countDay += 1
            sum = weight
        }
    }
    
    return countDay + 1
}
