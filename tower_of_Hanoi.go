package main

//三部曲
//1、把 n-1 号盘子移动到缓冲区
//2、把1号从起点移到终点
//3、然后把缓冲区的n-1号盘子也移到终点
func hanota(A []int, B []int, C []int) []int {
	if A == nil {
		return nil
	}
	var move func(int, *[]int, *[]int, *[]int)
	move = func(num int, from, buffer, to *[]int) {
		if num < 0 {
			return
		}
		if num == 1 {
			*to = append(*to, (*from)[len(*from)-1])
			*from = (*from)[:len(*from)-1]
			return
		}
		move(num-1, from, to, buffer)
		move(1, from, buffer, to)
		move(num-1, buffer, from, to)
	}
	move(len(A), &A, &B, &C)
	return C
}