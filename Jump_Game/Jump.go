package Jump_Game

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

//55. 跳跃游戏  贪心
//输入: [3,2,1,0,4]
//输出: false
func canJump(nums []int) bool {
	maxPotion := nums[0] //最大可达边界
	for i:= 0 ; i < len(nums) ; i++ {
		if i <= maxPotion {
			if i+nums[i] > maxPotion {
				maxPotion = i+nums[i]
			}
		}else{
			return false
		}
	}
	return true
}

//45. 跳跃游戏 II
func jump(nums []int) int {
	maxPotion := 0 //更新最大边界状态
	length := len(nums)-1
	step := 0
	end := 0 //最大边界
	for i := 0 ; i < length ; i++ {
		maxPotion = max(maxPotion,i+nums[i])
		if i == end {
			end = maxPotion
			step++
		}
	}
	return step
}

//1306. 跳跃游戏 III bfs
func canReach(arr []int, start int) bool {
	if arr[start] == 0 {
		return true
	}
	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true
	start = 0
	for start < len(queue) {
		id := queue[start]
		start++

		idx := id + arr[id]
		if idx >= 0 && idx < len(arr){
			if !visited[idx]{
				if arr[idx] == 0 {
					return true
				}
				queue = append(queue,idx)
				visited[idx] = true
			}
		}

		idx = id - arr[id]
		if idx >= 0 && idx < len(arr) {
			if !visited[idx] {
				if arr[idx] == 0{
					return true
				}
				queue = append(queue, idx)
				visited[idx] = true
			}
		}
	}
	return false
}

//1345. 跳跃游戏 IV


//1340. 跳跃游戏 V