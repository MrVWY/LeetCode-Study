package Array

//845. 数组中的最长山脉   36
func longestMountain(A []int) int {
	length := len(A)
	if length < 3 {
		return 0
	}
	ans, base := 0 , 0
	for base < length{
		end := base
		if end + 1 < length && A[end] < A[end + 1] {
			for end + 1 < length && A[end] < A[end + 1] {
				end++
			}

			if end + 1 < length && A[end] > A[end + 1] {
				for end + 1 < length && A[end] > A[end + 1] {
					end++
				}
				ans = maxs2(ans, end - base + 1)    // int(math.Max(float1,float2))    math.Max 要求是两个float
			}
		}
		base = maxs2(end , base + 1 )
	}
	return ans
}

func maxs2(a , b int) int {
	if a > b {
		return a
	}
	return b
}

//1052. 爱生气的书店老板
func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum, lost := 0,0
	for i := 0 ; i < len(customers) ; i++ {
		sum += customers[i]
		if grumpy[i] == 1 {
			lost += customers[i]
		}else {
			customers[i] = 0
		}
	}
	max , left , clam := 0,0,0
	for r := 0 ; r < len(customers) ; i++ {
		if r - left + 1 <= X{
			clam += customers[r]
		}else {
			clam += customers[r]
			clam -= customers[left]
			left++
		}
		if clam > max {
			max = clam
		}
	}
	return sum - lost + clam
}

//面试题 17.16. 按摩师   典型的动态规定
func massage(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp0 , dp1 := 0 , nums[0]
	for i := 1 ; i < length ; i++ {
		tmp0 := max(dp0,dp1)
		tmp1 := dp0 + nums[i]
		dp0 , dp1 = tmp0 , tmp1
	}
	return max(dp0 , dp1)
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

