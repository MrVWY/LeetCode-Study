package main

import (
	"fmt"
)

/*
			abc
         db   nm
      f   g      h
abcdbf  abcdbg   abcnmh
*/

type Node struct {
	Nodes []*Node
	Index string  //所有Nodes节点中字符串的首字母集合 当前节点下的路径字母
	Path string   //当前节点的字符串
}
var c = NewNode("root")

func NewNode(s string)  *Node{
	return &Node{
		Nodes:make([]*Node,0),
		Index:"",
		Path:s,
	}
}

func findSameI(a,b string)  int{
	i := 0
	max := min(len(a),len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	// 返回一个下标, 表面当前下标中的两个字符不同
	return i
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//核心方法
func (t *Node) Insert(p string)  {

	// 根节点下没有任何子节点
	if len(t.Nodes) == 0 {
		n := &Node{
			Nodes:make([]*Node,0),
			Index:"",
			Path:p,
		}
		t.Nodes = []*Node{n}
		t.Index = string([]byte{p[0]})   //root->oll
		fmt.Println("根节点下没有任何子节点", t.Path, t.Nodes[0].Path)
		return
	} else {
		//查找插入的字符串是否包含在当前节点的子节点中
		te := false
		for i,max := 0,len(t.Index);i<max; i++ {
			if p[0] == t.Index[i] {
				fmt.Println("p", t.Index)
				// 切换到相应的子节点下
				t = t.Nodes[i]
				fmt.Println("切换到相应的子节点下", t.Index, t.Path)
				te = true
				break
			}
		}
		// 如果没有相同的节点，就在当前跟节点下新开辟一个新的子节点
		if !te {
			n := &Node{
				Nodes:make([]*Node,0),
				Index:"",
				Path:p,
			}
			t.Nodes = append(t.Nodes,n)
			t.Index += string([]byte{p[0]})
			fmt.Println("如果没有相同的节点，就在当前跟节点下新开辟一个新的子节点", t.Index)
			return
		}
	}

walk:
	for {
		//查找相同的字符串的数字
		i := findSameI(t.Path,p)
		fmt.Println("查找相同的字符串的数字",i, t.Path, p)
		//如果相同的字符串与当前节点的字符串小，说明
		//需要帮当前节点进行拆分
		// abc   abd
		// 		a (path = ab)
		// 	  b
		//  c  (拆分)
		// 如果在第i位不一样，则需要帮其进行拆分
		if i < len(t.Path) {
			child := &Node{
				Nodes:t.Nodes,
				Path:t.Path[i:],
				Index:t.Index,
			}
			fmt.Println("before Index",t.Index, "Path",t.Path)
			fmt.Println("child path", child.Path, "child Index",child.Index)
			t.Index = string([]byte{t.Path[i]})
			t.Path = t.Path[:i]
			fmt.Println("after Index",t.Index, "Path",t.Path)
			t.Nodes = []*Node{child}
		}

		//如果相同的字符串与插入的字符串小，说明需要
		//创建新的节点
		if i <len(p) {
			p = p[i:]  //ollpm
			c := p[0]  //pm
			for i,max := 0,len(t.Index);i<max; i++ {
				if c == t.Index[i] {
					// 切换到下一个子节点
					t = t.Nodes[i]
					fmt.Println("t.",t.Index)
					continue walk
				}
			}
			// 不需要切换到下一个子节点，则按原路径进行增添路径
			child := &Node{
				Nodes:make([]*Node,0),
				Path:p,
				Index:"",
			}
			t.Index += string([]byte{c})
			fmt.Println("end", t.Index, child.Path)
			t.Nodes = append(t.Nodes,child)
		}
		return
	}
}

func main()  {
	c.Insert("ollid")
	c.Insert("ollpm")
	c.Insert("ollpmc")
	//c.Insert("liu")
	//c.Insert("lpmc")
	fmt.Println(c)
}