func hasAllCodes(s string, k int) bool {
 
    // input string size as well as code size
    strSize, codeLen := len(s), k
    
	// Go doesn't have native set, so we use map(i.e., so-called dictionary) as alternative plan
    codeDict := make( map[string]bool)
    
    // scan each substring with code length, save into code dictionary
    for start :=0 ; start < strSize - codeLen + 1 ; start += 1{
        
        cur_code := s[ start : start + codeLen ]
        
        if _, exist := codeDict[cur_code]; exist == false{
            codeDict[cur_code] = true
        }
    }
    
    // we have all binary codes if code set size = 2^k
    return len(codeDict) == int( math.Pow(2.0, float64(k) ) )
    
}
