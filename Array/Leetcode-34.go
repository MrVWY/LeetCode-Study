package Array

//289、生命游戏   重点理解 “同步更新” ！！！！！！！！！！！！！
func gameOfLife(board [][]int)  {
	temp := make([][]int,len(board))
	for i := 0 ; i < len(board) ; i++ {
		temp[i] = make([]int,len(board[0]))
	}

	for i := 0 ; i < len(board) ; i++{
		for j := 0 ; j < len(board[0]) ; j++ {
			nums := 0
			// 上面
			if i - 1 >= 0 {
				nums += board[i-1][j]
			}
			// 左面
			if j - 1 >= 0 {
				nums += board[i][j-1]
			}
			// 下面
			if i + 1 < len(board) {
				nums += board[i+1][j]
			}
			// 右面
			if j + 1 < len(board[i]) {
				nums += board[i][j+1]
			}
			// 左上
			if i - 1 >= 0 && j - 1 >= 0 {
				nums += board[i-1][j-1]
			}
			// 右上
			if i - 1 >= 0 && j + 1 < len(board[i]) {
				nums += board[i-1][j+1]
			}
			// 左下
			if i + 1 < len(board) && j - 1 >= 0  {
				nums += board[i+1][j-1]
			}
			// 右下
			if j + 1 < len(board[i]) && i + 1 < len(board) {
				nums += board[i+1][j+1]
			}
			temp[i][j] = board[i][j]
			switch {
			case nums < 2: temp[i][j] = 0
			case nums == 3 && temp[i][j] == 0: temp[i][j] = 1
			case nums > 3: temp[i][j] = 0
			}
		}
	}
	copy(board, temp)
}

//1162. 地图分析   0 代表海洋，1 代表陆地
type piont struct {
	X int
	Y int
}

func maxDistance(grid [][]int) int {
	dx := []int{-1,0,1,0}
	dy := []int{0,1,0,-1}

	row := len(grid)
	ver := len(grid[0])

	var queue []*piont

	for i, row := range grid  {
		for j, v := range row {
			if v == 1 {
				queue = append(queue,&piont{i,j})
			}
		}
	}

	if len(queue) == 0 || len(queue) == row * ver {
		return -1
	}

	count := 0
	for len(queue) > 0 {
		count++
		temp := queue
		queue = nil
		for len(temp) > 0 {
			p := temp[0]
			temp = temp[1:]
			for i := 0 ; i < 4 ; i++ {
				x := dx[i] + p.X
				y := dy[i] + p.Y
				if x < 0 || x >= row || y < 0 || y >= ver ||  grid[x][y] != 0 {
					continue
				}
				queue = append(queue, &piont{x, y})
				grid[x][y] = 2
			}
		}
	}
	return count - 1
}