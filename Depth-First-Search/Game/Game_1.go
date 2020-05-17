package Game

//529. 扫雷游戏
var d = [8][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
func updateBoard(board [][]byte, click []int) [][]byte {
	a, b := click[0], click[1]
	var f func(x, y int)
	f = func(i, j int) {
		count := byte('0')
		for _, v := range d {
			x, y := i+v[0], j+v[1]
			if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) && board[x][y] == 'M' {
				count++
			}
		}

		if count > '0' {
			board[i][j] = count
		}else{
			board[i][j] = 'B'
			for _, v := range d {
				x, y := i+v[0], j+v[1]
				if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) && board[x][y] == 'E' {
					f(x,y)
				}
			}
		}
	}

	if board[a][b] == 'M' {
		board[a][b] = 'X'
	}else if board[a][b] == 'E'{
		f(a,b)
	}

	return board
}