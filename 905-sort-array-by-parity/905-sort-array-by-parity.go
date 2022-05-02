func sortArrayByParity(nums []int) []int {
    evenIndex := 0
    oddIndex := len(nums) - 1
    
    for evenIndex < oddIndex {
         if nums[evenIndex] % 2 != 0 {
            nums[evenIndex], nums[oddIndex] = nums[oddIndex], nums[evenIndex]
            oddIndex --
         } else {
             evenIndex ++
         }
    }
    
    return nums
}