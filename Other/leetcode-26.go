package main

//191. 位1的个数 (与运算)
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		res++
		num &= (num - 1)
	}
	return res
}

//187. 重复的DNA序列
func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	rec := make(map[string]int,len(s)-9)
	if len(s) < 10 {
		return res
	}

	for i := 0 ; i+10<=len(s) ; i++ {
		c := s[i:i+10]
		if v, _ := rec[c]; v == 1 {
			res = append(res,c)
		}
		rec[c]++
	}
	return res
}


