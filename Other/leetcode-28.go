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

//直线上最多的点数
type Point struct {
	x, y int
}

type K struct {
	k1 , k2 int
}

type B struct {
	b1 , b2 int
}

type Line struct {
	K
	B
	c int
}

//求最小公约数(辗转相除法)
func gcd(a,b int) int {
	for b != 0 {
		a, b = b , a%b
	}
	return a
}

func NewLine(p,q Point) Line {
	if p.x == q.x {
		return Line{c: p.x}
	}

	k1 , k2 := p.y - q.y , p.x - q.x
	g1 := gcd(k1,k2)
	k := K{k1/g1,k2/g1}

	b1 , b2 := q.y * k.k2 - k.k1 * q.x , k.k2   // b1 / b2  =  (y * k2 - k1 * x1) / k2
	g2 := gcd(b1,b2)
	b := B{b1/g2,b2/g2}

	return Line{K: k, B: b}
}

func (L Line) HasPointInLine(p Point) bool {
	if L.k2 == 0 {
		//垂直X的直线
		return p.x == L.c
	}
	return p.y * L.k2 * L.b2 == L.k1 * L.b2 * p.x + L.k2 * L.b1  //y * k2 * b2 = k1 * b1 * x + k2 * b1
}

func maxPoints(points [][]int) int {
	length := len(points)
	if length < 3 {return length}

	i := 0
	for ; i < length - 1 ; i++ {
		if points[i][0] != points[i+1][0] || points[i][1] != points[i+1][1] {break}
	}
	if i == length - 1 { return length}

	lines := make(map[Line]int)
	var line Line
	count := 0
	for i := 0 ; i < length ; i++ {
		for j := i + 1 ; j < length ; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {continue} // 避免相同的点出现

			line = NewLine(Point{x: points[i][0], y: points[i][1]}, Point{x: points[j][0], y: points[j][1],})
			//查看line 是否在之前已经出现过
			if lines[line] != 0 {continue}

			count = 2
			//计算点是否在直线上
			for index := 0 ; index < length ; index++ {
				if index ==i || index ==j {continue}
				if line.HasPointInLine(Point{points[index][0], points[index][1]}) {count++}
			}

			//记录
			lines[line] = count
		}
	}
	max := 0
	for _ , v := range lines {
		if v > max {
			max = v
		}
	}
	return max
}