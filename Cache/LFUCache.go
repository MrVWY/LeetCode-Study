package Cache

//460
//核心部件Cache
type LFUCache struct {
	cache    map[int]*Node //节点表
	freq     map[int]*List //频率节点链表  x:frea-index y:Node
	//    y:
	//x:          		|-->  Node2 --->  Node3
	//   		Node1--|
	capacity int           //容量
	size     int           //当前大小 用于判断
	minFreq  int           //当前最小的频率
}
//节点
type Node struct {
	pre, Next *Node
	key, value, freq int // key-值-频率
}
//节点链表
//第0个为head,最后一个为tail
type List struct {
	head, tail *Node //头 尾
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:    make(map[int]*Node),
		freq:     make(map[int]*List),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	//判断当前节点表中是否含有key
	if node, ok := this.cache[key] ; ok{
		//如果含有key,频率+1,放到频率节点表的头部
		this.AddTheNodeFreq(node)
		return node.value
	}
	return -1
}


func (this *LFUCache) Put(key int, value int)  {
	if this.capacity == 0 {
		return
	}
	//判断当前节点表中是否含有key
	if node, ok := this.cache[key] ; ok {
		//注意key所对应的value是可以任意的
		node.value = value
		//频率+1,放到频率节点表的头部
		this.AddTheNodeFreq(node)
	}else {
		//判断当前Cache容量是否大于Cache容量
		if this.size >= this.capacity {
			node := this.freq[this.minFreq].RemoveLastNode()
			delete(this.cache, node.key)
			this.size--
		}
		NewNode := &Node{key: key, value: value, freq:1}
		this.cache[key] = NewNode
		if this.freq[1] == nil {
			this.freq[1] = GetTheList()
		}
		this.freq[1].AddFirst(NewNode)
		this.minFreq = 1
		this.size++
	}
}

//初始化一个List
func GetTheList() *List {
	head, tail := &Node{} , &Node{}
	head.Next, tail.pre = tail, head
	return &List{
		head: head,
		tail: tail,
	}
}

func (this *LFUCache) AddTheNodeFreq(node *Node) {
	_freq := node.freq //获取node原来的frea
	this.freq[_freq].RemoveNode(node)
	//判断最小频率minFreq是否等于_freq 和 node所在frea_index是否为空
	if this.minFreq == _freq && this.freq[_freq].IsEmpty() {
		this.minFreq++
		delete(this.freq, _freq)
	}
	node.freq++
	//判断node频率+1后所在的frea_index是否为空
	if this.freq[node.freq] == nil {
		this.freq[node.freq] = GetTheList()
	}
	this.freq[node.freq].AddFirst(node)
}

func (this *List) IsEmpty() bool {
	return this.head.Next == this.tail
}

func (this *List) RemoveNode(node *Node)  {
	node.pre.Next = node.Next
	node.Next.pre = node.pre

	node.pre = nil
	node.Next = nil
}

func (this *List) RemoveLastNode() *Node{
	if this.IsEmpty() {
		return nil
	}
	Last := this.tail.pre
	this.RemoveNode(Last)
	return Last
}

func (this *List) AddFirst(node *Node) {
	node.Next = this.head.Next
	node.pre = this.head

	this.head.Next.pre = node
	this.head.Next = node
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */