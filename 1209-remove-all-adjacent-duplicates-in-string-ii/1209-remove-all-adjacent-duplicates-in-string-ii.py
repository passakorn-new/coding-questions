class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        st = []  # pair of (character, frequence)
        for c in s:
            if st and st[-1][0] == c:
                st[-1][1] += 1  # Increase the frequency
            else:
                st.append([c, 1])
            if st[-1][1] == k:  # If reach enough k duplicate letters -> then remove
                st.pop()
        return "".join(c * f for c, f in st)