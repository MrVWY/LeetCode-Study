package Catch_rain

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//42、接雨水I  利用单调栈
func trap(height []int) int {
	sum := 0
	stack := make([]int,0)
	idx := 0
	for idx < len(height) {
		for len(stack) != 0 && height[idx] > height[stack[len(stack)-1]]{
			h := height[stack[len(stack)-1]]
			stack = stack[0:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			weight := idx - stack[len(stack)-1] -1
			min := min(height[stack[len(stack)-1]], height[idx])
			sum += weight * (min - h)
		}
		stack = append(stack,idx)
		idx++
	}
	return sum
}

//407. 接雨水 II