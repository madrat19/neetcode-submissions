type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	len     int
	cap     int
	first   *Node
	last    *Node
	storage map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		len:     0,
		cap:     capacity,
		storage: make(map[int]*Node),
		first:   &Node{},
		last:    &Node{},
	}
	cache.first.Next = cache.last
	cache.last.Prev = cache.first
	return cache
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.Prev, node.Next
	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) insert(node *Node) {
	prev, next := this.first, this.first.Next
	prev.Next = node
	node.Prev = prev
	node.Next = next
	next.Prev = node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.storage[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.insert(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.storage[key]
	if ok {
		this.remove(node)
		this.insert(node)
		node.Val = value
		delete(this.storage, key)
		this.storage[key] = node
	} else {
		node = &Node{Key: key, Val: value}
		this.insert(node)
		this.storage[key] = node
		this.len++
		if this.len > this.cap {
			node = this.last.Prev
			this.remove(node)
			delete(this.storage, node.Key)
		}
	}
}
