package main

//旋转图像
func rotatea(matrix [][]int)  {
	for i := range matrix {
		for j := range matrix[i] {
			if i < j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if j < len(matrix[i])/2 {
				matrix[i][j], matrix[i][len(matrix[i])-1-j] = matrix[i][len(matrix[i])-1-j], matrix[i][j]
			}
		}
	}
}

//字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	var v [26]int
	for _, c := range s {
		v[c-'a']++
	}
	for i,c := range s {
		if v[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

//有效的字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	var v [26]int
	for _, c := range s {
		v[c-'a']++
	}
	for _, c := range t {
		v[c-'a']--
	}
	for _, c := range s {
		if v[c-'a'] != 0 {
			return false
		}
	}
	return true
}