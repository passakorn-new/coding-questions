type NumMatrix struct {
	m [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	preRet := make([][]int, len(matrix)+1)
	for i := range preRet {
		preRet[i] = make([]int, len(matrix[0])+1)
	}
	for i := range matrix {
		for j := range matrix[i] {
			preRet[i+1][j+1] = preRet[i+1][j] + preRet[i][j+1] + matrix[i][j] - preRet[i][j]
		}
	}
	return NumMatrix{m: preRet}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.m[row2+1][col2+1] - this.m[row1][col2+1]  - this.m[row2+1][col1] + this.m[row1][col1]
}