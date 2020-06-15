package main

type LRUCache struct {
	Capacity int
	Cache    map[int]*Node2
	Head     *Node2
	Tail     *Node2
}

type Node2 struct {
	key   int
	value int
	Pre   *Node2
	Next  *Node2
}

func Constructor (capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*Node2),
		Head:     &Node2{},
		Tail:     &Node2{},
	}

	cache.Head.Next = cache.Tail
	cache.Tail.Pre = cache.Head
	return cache
}

//Set key and value to the Cache
func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Cache[key]
	if ok {
		this.removeNode(node)
	}else {
		if len(this.Cache) == this.Capacity {
			delete(this.Cache, this.Tail.Pre.key)
			this.removeNode(this.Tail.Pre)
		}
		node = &Node2{key: key, value:    value,}
		this.Cache[node.key] = node
	}
	node.value = value
	this.addNodeToHead(node)
}

//Get returns the values of the key
func (this *LRUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}
	this.removeNode(node)
	this.addNodeToHead(node)
	return node.value
}

func (this *LRUCache)removeNode(node *Node2)  {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache)addNodeToHead(node *Node2)  {
	this.Head.Next.Pre = node
	node.Next = this.Head.Next
	this.Head.Next = node
	node.Pre = this.Head
}