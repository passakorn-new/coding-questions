class Solution:
    def getSmallestString(self, n: int, k: int) -> str:
        remain_k = k - n                        # assume if we already have a (n times) -> (k - n)
        z_count = remain_k // 25                # divide 25 for find equal ascii other charactor
        other_ascii = remain_k % 25             # float 25 is ascii other charactor
        other_count = 1 if other_ascii else 0
        
        first_str = 'a' * (n - z_count - other_count)
        second_str = chr(97 + other_ascii) if other_ascii else ''
        last_str = 'z' * z_count
        return first_str + second_str + last_str
        
        # NOTES
        # result = a * amount_a + other_char + z * amount_z
