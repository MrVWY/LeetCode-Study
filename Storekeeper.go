package main

//推箱子
var (
	move = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, 1}, []int{0, -1}}
	n, m int
)

type node struct {
	x, y int
	step int
	peopleIndex [2]int
}

func minPushBox(grid [][]byte) int {
	if len(grid) == 0 {
		return -1
	}
	var people, box [2]int
	n, m = len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch grid[i][j] {
			case 'S':
				people = [2]int{i, j}
			case 'B':
				box = [2]int{i, j}
			}
		}
	}
	return bfs(grid, people, box)
}

func bfs(grid [][]byte, people, box [2]int) int {
	stack := make([]node, 1000)
	visited := make(map[[4]int]bool) //people and box
	start, end := 0, 0
	stack[start] = node {
		x:      box[0],
		y:      box[1],
		peopleIndex: people,
	}
	for start <= end {
		boxX, boxY, boxS := stack[start].x, stack[start].y, stack[start].step
		if grid[boxX][boxY] == 'T' {
			return boxS
		}
		peopleX, peopleY := stack[start].peopleIndex[0], stack[start].peopleIndex[1]
		start++
		for _, v := range move {
			boxmx, boxmy := boxX+v[0], boxY+v[1]
			if boxmx < 0 || boxmx >= n || boxmy < 0 || boxmy >= m {
				continue
			}
			if grid[boxmx][boxmy] == '#'{
				continue
			}
			peoplemx, peoplemy := boxX-v[0], boxY-v[1]
			if peoplemx < 0 || peoplemx >= n || peoplemy < 0 || peoplemy >= m {
				continue
			}
			if visited[[4]int{boxmx, boxmy, peoplemx, peoplemy}] {
				continue
			}
			if dfs(grid, map[[2]int]bool{[2]int{boxX,boxY}:true}, peopleX, peopleY, [2]int{boxX, boxY}, [2]int{peoplemx, peoplemy}) {
				if grid[boxmx][boxmy] == 'T' {
					return boxS+1
				}
				end++
				stack[end] = node{
					x:      boxmx,
					y:      boxmy,
					step:   boxS + 1,
					peopleIndex: [2]int{peoplemx, peoplemy},
				}
				//移动之后
				visited[[4]int{boxmx, boxmy, peoplemx, peoplemy}] = true
			}
		}
	}
	return -1
}
//看人是否可以到达对应的点
func dfs(grid [][]byte, visited map[[2]int]bool, peopleX, peopleY int, box, target [2]int) bool {
	visited[[2]int{peopleX, peopleY}] = true
	for _, v := range move {
		peoplemx, peoplemy := peopleX+v[0], peopleY+v[1]
		if peoplemx < 0 || peoplemx >= n || peoplemy < 0 || peoplemy >= m {
			continue
		}
		if visited[[2]int{peoplemx, peoplemy}] || grid[peoplemx][peoplemy] == '#' || (peoplemx == box[0] && peoplemy == box[1]){
			continue
		}
		if peoplemx == target[0] && peoplemy == target[1]{
			return true
		}
		if dfs(grid, visited, peoplemx, peoplemy, box, target) {
			return true
		}
	}
	return false
}