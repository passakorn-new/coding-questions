func findUnsortedSubarray(nums []int) int {
    lenNums := len(nums)
    min := nums[lenNums - 1]
    max := nums[0]
    begin, end := -1, -2
    
    for i := 1; i < lenNums; i++ {
        if max < nums[i] {
            max = nums[i]
        }
        
        if min > nums[lenNums - 1 - i] {
            min = nums[lenNums - 1 - i]
        }
        
        if nums[i] < max { end = i }
        if nums[lenNums - 1 - i] > min { begin = lenNums - 1 - i }
    }
    
    return end - begin + 1;
}