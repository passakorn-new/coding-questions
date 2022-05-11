func countVowelStrings(n int) int {
    dp := [5]int{1,1,1,1,1}

	for i := 0; i < n-1; i++ {
        next := [5]int{}

		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				next[j] += dp[k]
			}

		}

		dp = next
	}

	var count int
	for i := range dp {
		count += dp[i]
	}
	return count
}