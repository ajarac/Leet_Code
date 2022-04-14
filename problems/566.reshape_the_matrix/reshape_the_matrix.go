package reshape_the_matrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	totalSize := r * c
	matTotalSize := len(mat) * len(mat[0])
	if totalSize != matTotalSize {
		return mat
	}
	newMat := make([][]int, r)
	cols := len(mat[0])
	var cRow, cCol int
	for indexR := 0; indexR < r; indexR++ {
		newMat[indexR] = make([]int, c)
		for indexC := 0; indexC < c; indexC++ {
			newMat[indexR][indexC] = mat[cRow][cCol]
			cCol++
			if cCol >= cols {
				cCol = 0
				cRow++
			}
		}
	}
	return newMat
}
