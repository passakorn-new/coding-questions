class Solution:
    def calPoints(self, ops: List[str]) -> int:
        current_score = []
        for i in range(0, len(ops)):
            if ops[i] == '+':
                current_score.append(current_score[-1] + current_score[-2])
            elif ops[i] == 'D':
                current_score.append(current_score[-1] * 2)
            elif ops[i] == 'C':
                current_score.pop(-1)
            else:
                current_score.append(int(ops[i]))

        return sum(current_score)
    
        # Notes
        # array[-1] => last element
        # array[-2] => 2th last element