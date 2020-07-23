package main

import "fmt"

//计算器 逆波兰表达式
func getPriority(ch byte) int {
	priority := 0
	switch ch {
		case '+': priority = 1
		case '-': priority = 1
		case '(': priority = 2
		case '/': priority = 3
		case '*': priority = 3
		default:
			priority = 0
	}
	return priority
}

func changeRPN(ch string) string {
	nums := make([]byte,0) //store number
	stack := make([]byte,0) //store operator
	for i := 0 ; i < len(ch) ; i++ {
		if ch[i] == ' ' {continue} //分隔符
		if ch[i] >= '0' && ch[i] <= '9' {
			//如果是二位数以上
			for i < len(ch) && ch[i] >= '0' && ch[i] <= '9' {
				nums = append(nums, ch[i])
				i++
			}
			nums = append(nums, ' ') //分隔符,不然会出现11*2 -> 112*
			i--
		}

		if ch[i] == '(' {
			//直接入栈
			stack = append(stack, ch[i])
		}else if ch[i] == ')' {
			//将括号中间的符号输出到尾部
			for len(stack) != 0 {
				top := stack[len(stack)-1]
				if top == '(' {
					stack = stack[:len(stack)-1]
					break
				}
				nums = append(nums, top)
				nums = append(nums, ' ')
				stack = stack[:len(stack)-1]
			}
		}

		if ch[i] == '+' || ch[i] == '-' || ch[i] == '*' || ch[i] == '/' {
			if len(stack) == 0 {
				stack = append(stack, ch[i])
			}else {
				if getPriority(ch[i]) > getPriority(stack[len(stack)-1]) {
					stack = append(stack, ch[i])
				}else {
					//如果当前运算符优先级比运算栈栈顶的运算符优先级小,将栈顶运算符输出到尾部
					for len(stack) != 0 && getPriority(ch[i]) <= getPriority(stack[len(stack)-1]) && stack[len(stack)-1] != '(' {
						nums = append(nums, stack[len(stack)-1])
						nums = append(nums, ' ')
						stack = stack[:len(stack)-1]
					}
					stack = append(stack, ch[i])
				}
			}
		}
	}
	fmt.Println(stack)
	for len(stack) != 0 {
		nums = append(nums, stack[len(stack)-1])
		nums = append(nums, ' ')
		stack = stack[:len(stack)-1]
	}

	result := ""
	for i := 0 ; i < len(nums) ; i++ {
		result += string(nums[i])
	}
	return result
}

func calculateRPN(ch string) int {
	length := len(ch)
	var value1, value2, result int
	stack := make([]int,0)
	for i := 0 ; i < length ; i++ {
		if ch[i] >= '0' && ch[i] <= '9' {
			tmp := int(ch[i] - 48)
			j := i+1
			for ch[j] >= '0' && ch[j] <= '9' {
				tmp = tmp * 10 + int(ch[j] - 48)
				j++
			}
			stack = append(stack, tmp)
			i = j-1
		}
		if ch[i] == '+' || ch[i] == '-' || ch[i] == '*' || ch[i] == '/' {
			value2 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			value1 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ch[i] == '+' { result = value1+value2
			}else if ch[i] == '-' {result = value1-value2
			}else if ch[i] == '*' {result = value1*value2
			}else if ch[i] == '/' {result = value1/value2}
			stack = append(stack, result)
		}
	}
	return stack[len(stack)-1]
}


func main() {
	ch := "( 12 * 12 + 3 ) * 2"
	fmt.Println(ch)
	fmt.Println(calculateRPN(changeRPN(ch)))
}