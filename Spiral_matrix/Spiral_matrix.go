package Spiral_matrix

//54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	res := make([]int,0)
	if len(matrix) == 0 {return res}
	weight, height := len(matrix[0]), len(matrix)
	up, down := 0, height-1
	left, right := 0, weight-1
	total := weight * height
	for {
		for i := left ; i <= right ; i++ {
			res = append(res, matrix[up][i])
		}
		if len(res) == total {
			break
		}

		up++
		for i := up ; i <= down ; i++ {
			res = append(res, matrix[i][right])
		}
		if len(res) == total {
			break
		}

		right--
		for i := right ; i >= left ; i-- {
			res = append(res, matrix[down][i])
		}
		if len(res) == total {
			break
		}

		down--
		for i := down ; i >= up ; i-- {
			res = append(res, matrix[i][left])
		}
		if len(res) == total {
			break
		}
		left++
	}

	return res
}

//59. 螺旋矩阵 II
func generateMatrix(n int) [][]int {
	res := make([][]int,n)
	for i := 0 ; i < len(res) ; i++ {
		res[i] = make([]int,n)
	}
	up, down := 0, n-1
	left, right := 0, n-1
	num, total := 1 , n*n
	for num <= total {
		for i := left ; i <= right ; i++ {
			res[up][i] = num
			num++
		}
		up++
		for i := up ; i <= down ; i++ {
			res[i][right] = num
			num++
		}
		right--
		for i := right ; i >= left ; i-- {
			res[down][i] = num
			num++
		}
		down--
		for i := down ; i >= up ; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}

//885. 螺旋矩阵 III
