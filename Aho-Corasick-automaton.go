package main

//AC自动机
//前置知识：前缀树+KMP

type Aho struct {
	next [26]int
	fail, cnt int //失配指针、节点信息
}
const maxLength = 1000
func newAho() [maxLength]*Aho {
	tree := [maxLength]*Aho{}
	root := &Aho{
		next: [26]int{},
		fail : 0,
		cnt : 0,
	}
	tree[0] = root
	return tree
}

var size = 1
//注意insert函数中可能会出现下标越界,待解决
//tree初始化是为0
func insert(s string, tree []*Aho) {
	n := len(s)
	now := 0
	for i := 0 ; i < n ; i++ {
		c := s[i]
		if tree[now].next[c-'a'] == 0 {
			tree[now].next[c-'a'] = size
			size++
		}
		now = tree[now].next[c-'a']
	}
	tree[now].cnt++
}

func build(tree []*Aho) {
	//宽搜 处理失配指针
	tree[0].fail = -1
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) != 0 {
		father := queue[0]
		queue = queue[1:]
		for i := 0 ; i < 26 ; i++ {
			//i代表一个字母(先声明,免得下去就乱了)
			//寻找father节点下的子节点
			if tree[father].next[i] != 0 {
				//如果father为根节点root
				if father == 0 {
					tree[tree[father].next[i]].fail = 0
				}else {
					v := tree[father].fail //father的失配指针
					//向上寻找
					for v != -1 {
						//father失配指针下是否有对应的节点
						if tree[v].next[i] != 0 {
							tree[tree[father].next[i]].fail = tree[v].next[i]
						}
						//如果father失配指针没有对应的节点
						v = tree[v].fail
					}
					//真就找不到,只能指向根节点
					if v == -1 {
						tree[tree[father].next[i]].fail = 0
					}
				}
				//加入宽搜队列
				queue = append(queue, tree[father].next[i])
			}
		}
	}
}

func match(s string, tree []*Aho) int {
	n := len(s)
	res, now := 0, 0
	for i := 0 ; i < n ; i++ {
		c := s[i]
		if tree[now].next[c-'a'] != 0 {
			now = tree[now].next[c-'a']
		}else {
			//匹配不上就只能转去失配指针
			point := tree[now].fail
			//如果point = -1 那么说明回到了根节点-最顶上的那个节点
			for point != -1 && tree[point].next[c-'a'] == 0 {
				point = tree[point].fail
			}
			if point == -1 {
				now = 0
			}else {
				now = tree[point].next[c-'a']
			}
		}
		//计算结果?????
		for now != 0 {
			res += tree[now].cnt
			tree[now].cnt = 0
			now = tree[now].fail
		}
	}
	return res
}