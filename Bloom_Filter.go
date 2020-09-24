package main

type Node struct {
	ch [26]int //字母-当前长度
	flag int   //下一个节点
}

var tree []Node

func Insert(word string, id int){
	add := 0
	for i := 0 ; i < len(word) ; i++ {
		str := int(word[i]-'a')
		if tree[add].ch[str] == 0 {
			tree = append(tree, Node{[26]int{}, -1})
			tree[add].ch[str] = len(tree)-1
		}
		add = tree[add].ch[str]
	}
	tree[add].flag = id
}

func findWord(word string, left, right int) int {
	add := 0
	for i := right ; i >= left ; i-- {
		str := int(word[i]-'a')
		if tree[add].ch[str] == 0 {
			return -1
		}
		add = tree[add].ch[str]
	}
	return tree[add].flag
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right - left + 1) / 2; i++ {
		if s[left + i] != s[right - i] {
			return false
		}
	}
	return true
}

func palindromePairs(words []string) [][]int {
	tree = []Node{Node{[26]int{}, -1}}
	n := len(words)
	for i := 0; i < n; i++ {
		Insert(words[i], i)
	}
	ret := [][]int{}
	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		for j := 0; j <= m; j++ {
			if isPalindrome(word, j, m - 1) {
				leftId := findWord(word, 0, j - 1)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word, 0, j - 1) {
				rightId := findWord(word, j, m - 1)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}