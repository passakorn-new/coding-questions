class Solution:
    def minDominoRotations(self, tops: List[int], bottoms: List[int]) -> int:
        list_size = len(tops)
        freq_tops = [0] * 7
        freq_bottoms = [0] * 7
        equal_indexes = [0] * 7

        for index in range(list_size):
            freq_tops[tops[index]] += 1
            freq_bottoms[bottoms[index]] += 1   
            
            if tops[index] == bottoms[index]:
                equal_indexes[tops[index]] += 1

        for index in range(7):
            if freq_tops[index] + freq_bottoms[index] - equal_indexes[index] == list_size:
                return list_size - max(freq_tops[index], freq_bottoms[index])
        return  -1
            
        # NOTES
        # freq_tops = Keep freq value from tops
        # freq_bottoms = Keep freq value from freq_bottoms
        # equal_indexes = Keep freq at same value
