package Design

//706. 设计哈希映射  拉链法
type MyHashMap struct {
	content map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	my := MyHashMap{}
	my.content = map[int]int{}
	return my
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	this.content[key] = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	val, ok := this.content[key]
	if !ok {
		return -1
	}
	return val
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	delete(this.content, key)
}