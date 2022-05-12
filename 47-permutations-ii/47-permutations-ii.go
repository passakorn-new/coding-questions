func permuteUnique(nums []int) [][]int {
    res := [][]int{}
    perm := []int{}
    count := make(map[int]int, 0)
    
    for _, n := range nums {
        count[n]++
    }
    
    backtracking(&res, perm, nums, count)
    return res
}

func backtracking(res *[][]int, perm []int, nums []int, count map[int]int)  {
    if len(perm) == len(nums) {
        //copy list into temp, so it won't impact subsequent process on list
        temp := make([]int, len(perm))
        copy(temp, perm)
        
        *res = append(*res, temp) 
        return
    }
    
    for n, v := range count {
        if v > 0 {
            perm = append(perm, n)
            count[n]--
            backtracking(res, perm, nums, count)
            perm = perm[:len(perm)-1]
            count[n]++
        }
    }    
}