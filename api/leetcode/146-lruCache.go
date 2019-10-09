package leetcode

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Hard (24.65%)
 * Total Accepted:    282.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 *
 * Design and implement a data structure for Least Recently Used (LRU) cache.
 * It should support the following operations: get and put.
 *
 *
 *
 * get(key) - Get the value (will always be positive) of the key if the key
 * exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present.
 * When the cache reached its capacity, it should invalidate the least recently
 * used item before inserting a new item.
 *
 *
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 *
 * Example:
 *
 * LRUCache cache = new LRUCache( 2 /* capacity  );
*
* cache.put(1, 1);
* cache.put(2, 2);
* cache.get(1);       // returns 1
* cache.put(3, 3);    // evicts key 2
* cache.get(2);       // returns -1 (not found)
* cache.put(4, 4);    // evicts key 1
* cache.get(1);       // returns -1 (not found)
* cache.get(3);       // returns 3
* cache.get(4);       // returns 4
*
*
*/
type LRUCache struct {
	head *Node
	tail *Node
	kv   map[int]*Node
	cap  int
}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

//func Constructor(capacity int) LRUCache {
//	return LRUCache{kv: make(map[int]*Node, 0), cap: capacity}
//}

func (this *LRUCache) Get(key int) int {
	var val int
	node, ok := this.kv[key]
	switch ok {
	case true:
		val = node.val
		this.removeNode(node)
		this.setHead(node)
	default:
		val = -1
	}
	return val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.kv[key]
	switch ok {
	case true:
		node.val = value
		this.removeNode(node)
		this.setHead(node)
	default:
		if len(this.kv) >= this.cap {
			delete(this.kv, this.tail.key)
			this.removeNode(this.tail)
		}
		newNode := &Node{key: key, val: value}
		this.setHead(newNode)
		this.kv[key] = newNode
	}
}

func (this *LRUCache) setHead(node *Node) {
	if this.head == nil {
		this.head = node
		this.tail = this.head
		return
	}

	node.next = this.head
	node.prev = nil
	this.head.prev = node
	this.head = node
}

func (this *LRUCache) removeNode(node *Node) {
	switch node.prev == nil {
	case true:
		this.head = node.next
	default:
		node.prev.next = node.next
	}

	switch node.next == nil {
	case true:
		this.tail = node.prev
	default:
		node.next.prev = node.prev
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
