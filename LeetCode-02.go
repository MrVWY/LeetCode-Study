package main

import "fmt"

//LRU
//双向链表                                                 	⬆---------------------------⬇
//			 ⬆----------------⬇	     ⬆----------------⬇   	⬆					       null
//     head(Next)     (Prev)Node_1(Next)     (Prev)Node_2(Next)       (Prev)Tail
//null                   ⬇     ⬆----------------⬇      ⬆-----------------⬇
//	⬆--------------------⬇
type LRUCache struct {
	Capacity int //容量
	Len int      //长度
	Map map[int]*Node
	head *Node   //双向链表头部
	Tail *Node   //双向链表尾部
}

type Node struct {
	Prev *Node  //Node上一个
	Next *Node  //Node下一个
	Value int
	Key int
}

func NewConstruction(capacity int) LRUCache {
	m := make(map[int]*Node)
	LRU := LRUCache{Capacity: capacity, Map:m, head:&Node{}, Tail:&Node{}}
	LRU.head.Next = LRU.Tail
	LRU.Tail.Prev = LRU.head
	return LRU
}

func (L *LRUCache) Get(key int) int {
	if v, ok := L.Map[key]; ok{
		//将v从双向链表抽离出来，同时v的前后俩个Node要连上
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		//将V插进双向链表的头部
		n := L.head.Next
		L.head.Next = v
		v.Prev = L.head
		n.Prev = v
		v.Next = n

		return v.Value
	}
	return -1
}

func (L *LRUCache) Put(key int, value int) {
	if v, ok := L.Map[key]; ok{
		//将v从双向链表抽离出来，同时v的前后俩个Node要连上
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		//将V插进双向链表的头部
		n := L.head.Next
		L.head.Next = v
		v.Prev = L.head
		n.Prev = v
		v.Next = n
		return
	}
	//如果双向链表没有超出容量的限制
	if L.Len < L.Capacity {
		L.Len ++
		NewNode := &Node{
			Value: value,
			Key:   key,
		}
		L.Map[key] = NewNode
		n := L.head.Next
		L.head.Next = NewNode
		NewNode.Prev = L.head
		n.Prev = NewNode
		NewNode.Next = n
	}else {
		t := L.Tail.Prev
		L.Tail.Prev.Prev.Next = L.Tail
		L.Tail.Prev = L.Tail.Prev.Prev
		//将新值赋给t
		t.Value = value
		//删除旧的
		delete(L.Map,t.Key)
		//将新值赋给t
		t.Key = key
		//将新值插到双向链表头部
		L.Map[key] = t
		n := L.head.Next
		L.head.Next = t
		t.Prev = L.head
		n.Prev = t
		t.Next = n
	}
	return
}

func main()  {
	Link := NewConstruction(3)
	Link.Put(1,3)
	Link.Put(2,4)
	Link.Put(3,5)
	Link.Put(4,6)
	a := Link.Get(1)
	fmt.Println(a)
}