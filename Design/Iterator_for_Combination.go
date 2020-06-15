package Design

import "fmt"

//1286. 字母组合迭代器  可以利用深度优先遍历DFS
type CombinationIterator struct {
	characters string
	flag bool
	data []int
}

//type CombinationIterator struct { paths []string }  先求出所有可能的组合  提前算出来放在数组里
func dfs(s string, length, index int, path string, paths *[]string) {
	if len(path) == length {
		*paths = append(*paths, path)

		return
	}

	for i := index; i < len(s); i++ {
		dfs(s, length, i + 1, fmt.Sprintf("%s%c", path, s[i]), paths)
	}
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	data := make([]int,combinationLength)
	for i := 0 ; i < combinationLength ; i++ {
		data[i] = i
	}
	return CombinationIterator{
		characters : characters,
		flag : false,
		data : data,
	}
}


func (this *CombinationIterator) Next() string {
	result := ""
	i := -1
	for _,v := range(this.data) {
		result = fmt.Sprintf("%s%c", result, this.characters[v])
	}
	for r := len(this.data) - 1 ; r > -1 ; r-- {
		if this.data[r] != len(this.characters) - len(this.data) + r {
			i = r
			break
		}
	}

	if i == -1 {
		this.flag = true
	}else {
		this.data[i]++
		for j := i + 1 ; j < len(this.data) ; j++ {
			this.data[j] = this.data[j-1] + 1
		}
	}
	return result
}


func (this *CombinationIterator) HasNext() bool {
	return !this.flag
}
/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */