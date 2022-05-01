func backspaceCompare(s string, t string) bool {
    var ptr1, ptr2 int = len(s)-1, len(t)-1
    for ptr1 >= 0 || ptr2 >= 0 {
        
        // Back the first pointer
        // until we have no backspace
        // on the current char of s
        for ptr1 >= 0 && s[ptr1] == '#' {
            charToDelete := 1
            ptr1--
            for ptr1 >= 0 && charToDelete > 0 {
                if s[ptr1] == '#' {
                    charToDelete++
                } else {
                    charToDelete--
                }
                ptr1--
            }
        }
        
        // Back the second pointer
        // until we have no backspace
        // on the current char of t
        for ptr2 >= 0 && t[ptr2] == '#' {
            charToDelete := 1
            ptr2--
            for ptr2 >= 0 && charToDelete > 0 {
                if t[ptr2] == '#' {
                    charToDelete++
                } else {
                    charToDelete--
                }
                ptr2--
            }
        }

        // Boths currents chars of ptr1 (s) and ptr2 (t)
        // are not backspace (#), so we compare them
        if (ptr1 >= 0 && ptr2 < 0) ||
            (ptr1 < 0 && ptr2 >= 0) ||
            (ptr1 >= 0 && ptr2 >= 0 && s[ptr1] != t[ptr2]) {
            return false
        }
        ptr1--
        ptr2--
    }

    return ptr1 <= 0 && ptr2 <= 0
}