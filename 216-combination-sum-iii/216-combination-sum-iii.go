func combinationSum3(k int, n int) [][]int {
	return combSum3(k, n, 10)
}
// use bar to record the last trace value, only smaller ones would be visited next
func combSum3(k int, n int, bar int) [][]int {
	ret := make([][]int, 0)
	if k < 1 || n < 1 || bar < 1 {
		return ret
	}
	// if it arrives final part
	if k == 1 {
		if n < bar {
			return [][]int{{n}}
		} else {
			return ret
		}
	}
	// we need to go thought all possible trace
	for i := minInt(bar-1, n); i > 0; i-- {
		t := combSum3(k-1, n-i, i)
		for _, each := range t {
		// and record the cases in ret
			ret = append(ret, append(each, i))
		}
	}
	return ret
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}