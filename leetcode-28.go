package main

import "container/list"

//848. 字母移位
//golang中string底层是通过byte数组实现的.  中文字符在unicode下占2个字节,在utf-8编码下占3个字节,而golang默认编码正好是utf-8.
//rune 等同于int32,常用来处理unicode或utf-8字符
//byte 等同于int8，常用来处理ascii字符
func shiftingLetters52(S string, shifts []int) string {
	s := []byte(S)
	shifts = append(shifts,0)
	for i :=len(S) - 1 ; i>= 0; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
		s[i] = 'a' + (s[i] - 'a' + byte(shifts[i])) % 26
	}
	return string(s)
}

func shiftingLetters28(S string, shifts []int) string {
	var sum int
	for i := len(shifts)-1; i >= 0; i-- {
		sum = (sum+shifts[i]%26)%26
		shifts[i] = sum
	}
	s := make([]byte, len(S))
	for i := range S {
		mid := S[i]+byte(shifts[i])
		if mid > 'z' {
			s[i] = mid-'z'-1+'a'
		} else {
			s[i] = mid
		}
	}
	return string(s)
}

//933. 最近的请求次数
type RecentCounter struct {
	Lists *list.List
}


func Constructor2() RecentCounter {
	return RecentCounter{
		Lists : list.New(),
	}
}


func (this *RecentCounter) Ping(t int) int {
	this.Lists.PushBack(t)
	for this.Lists.Front().Value.(int) < (t - 3000) {
		this.Lists.Remove(this.Lists.Front())
	}
	return this.Lists.Len()
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

//面试题05. 替换空格
func replaceSpace(s string) string {
	result := ""
	for _ , v := range s {
		switch v {
		case ' ':
			result += "%20"
		default:
			result += string(v)
		}
	}
	return result
}

//1343. 大小为 K 且平均值大于等于阈值的子数组数目
func numOfSubarrays(arr []int, k int, threshold int) int {
	Left , Right := 0 , k - 1
	count , sum := 0, 0
	for i := 0 ; i <= Right ; i++ {
		sum += arr[i]
	}

	for Right != len(arr) - 1 {
		if sum >= k * threshold {
			count++
		}
		sum -= arr[Left]
		Left++
		Right++
		sum += arr[Right]
	}
	if sum >= k * threshold {
		count++
	}
	return count
}