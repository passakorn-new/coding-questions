func missingNumber(nums []int) int {
    var arrsum, allsum int
    for i:=0; i<len(nums); i++{
        arrsum+=nums[i]
        allsum+=i
    }
    return (allsum+len(nums))-arrsum
}