

func maxOperations(nums []int, k int) int {
    hashMap := make(map[int]int)
    res := 0
    
    for i := 0; i < len(nums); i++ {
        if hashMap[k - nums[i]] > 0 {
            res ++
            hashMap[k - nums[i]] --
        } else {
            hashMap[nums[i]] ++          
        }
    }
    
    return res
}