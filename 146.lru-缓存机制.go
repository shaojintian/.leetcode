/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	m        map[int]*DoubleLinkedListNode
	root     *DoubleLinkedListNode
	end      *DoubleLinkedListNode
	capacity int
	size     int
}

type DoubleLinkedListNode struct {
	key   int
	value int
	pre   *DoubleLinkedListNode
	next  *DoubleLinkedListNode
}

func Constructor(capacity int) LRUCache {
	m := make(map[int]*DoubleLinkedListNode, 0)
	root := &DoubleLinkedListNode{
		value: -1,
	}
	end := &DoubleLinkedListNode{
		value: -1,
	}
	root.next = end
	end.pre = root
	return LRUCache{
		m:        m,
		root:     root,
		end:      end,
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	//存在
	if v, ok := this.m[key]; ok {
		//refresh ddl
		this.removeNode(v)
		this.addToHead(v)
		return v.value
	}
	//不存在
	return -1
}

func initNode(key, value int) *DoubleLinkedListNode {
	return &DoubleLinkedListNode{
		key:   key,
		value: value,
	}
}

func (this *LRUCache) Put(key int, value int) {
	//存在
	if v, ok := this.m[key]; ok {
		v.value = value
		//refresh dll
		this.removeNode(v)
		this.addToHead(v)
	} else {
		//不存在且size小于capacity
		node := initNode(key, value)
		if this.size < this.capacity {
			this.m[key] = node
			this.addToHead(node)
			this.size++
		} else {
			//不存在且size == capacity,删除最久未使用
			this.m[key] = node
			key := this.removeTail()
			delete(this.m, key)
			this.addToHead(node)
		}
	}
}

func (this *LRUCache) removeNode(node *DoubleLinkedListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addToHead(node *DoubleLinkedListNode) {
	oldHead := this.root.next
	this.root.next = node
	node.pre = this.root
	node.next = oldHead
	oldHead.pre = node
}
func (this *LRUCache) removeTail() int {
	oldTailKey := this.end.pre.key
	tail := this.end.pre.pre
	tail.next = this.end
	this.end.pre = tail

	return oldTailKey
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

