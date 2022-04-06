class Solution:
    def threeSumMulti(self, arr: List[int], target: int) -> int:
        mod = (10 ** 9) + 7
        dict_map = {}
        res = 0
        
        for i in range(0, len(arr)):
            res = (res + dict_map.get(target - arr[i], 0)) % mod
            
            for j in range(0, i):
                two_sum = arr[i] + arr[j];
                dict_map[two_sum] = dict_map.get(two_sum, 0) + 1
        
        return res