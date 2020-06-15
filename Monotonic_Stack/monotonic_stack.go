package Monotonic_Stack

//739. 每日温度
func dailyTemperatures(T []int) []int {
	length := len(T)
	ans := make([]int, length)
	stack := []int{}
	for i := 0 ; i < length ; i++ {
		temperature := T[i]
		//当前元素(未入栈)与栈顶坐标所对应的值比较
		for len(stack) > 0 && temperature > T[stack[len(stack)-1]]{
			prevIndex := stack[len(stack)-1] //栈顶元素index
			stack = stack[:len(stack)-1] //移除栈顶坐标元素
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i) //将坐标压入栈中
	}
	return ans
}