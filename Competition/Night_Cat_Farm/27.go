package Night_Cat_Farm

//5410. 课程安排 IV
//打表法
//j -> i -> k 有向图路径
func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	m := make([][]bool, n)

	for i:=0 ; i <len(m) ; i++ {
		m[i] = make([]bool, n)
	}

	for _, v := range prerequisites {
		m[v[0]][v[1]] = true
	}

	for i := 0 ; i < n ; i++ {
		for j := 0 ; j < n ; j++ {
			for k := 0 ; k < n ; k++ {
				if m[j][i] && m[i][k] {
					m[j][k] = true
				}
			}
		}
	}

	res := make([]bool, len(queries))
	for k,v := range queries {
		res[k] = m[v[0]][v[1]]
	}
	return res
}