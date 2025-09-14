package main

import (
	"fmt"
	"sync"
)

// Node represents a node in a doubly linked list
type Node[K comparable, V any] struct {
	key K 
	value V 
	prev *Node[K, V]
	next *Node[K, V]
}


type LRUCache[K comparable, V any] struct {
	capacity int 
	head *Node[K, V]
	tail *Node[K, V]
	cache map[K]*Node[K, V]
	mu sync.Mutex

	/*
	head -> node1 -> node2-> node3-> node4 -> ..... -> tail 
	*/
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	if capacity <= 0 {
		panic("lru: capacity must be greater than 0")
	}

	head := &Node[K, V]{}
	tail := &Node[K, V]{}
	// set the head's next to tail
	head.next = tail 
	// Set the tail's prev to head 
	tail.prev = head 
	cache := make(map[K]*Node[K, V])
	return &LRUCache[K, V]{
		capacity: capacity,
		head: head, 
		tail: tail,
		cache: cache, 
	}
}

// Returns a key from the cache if it exists
// Time: O(1)
func(l *LRUCache[K, V]) Get(key K) (V , bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	node, ok := l.cache[key] 
	if !ok {
		var zero V
		return zero, false 
	}
	// remove this node from doubly linked list 
	l.removeNode(node)
	// Add this node to the head of the list
	l.addToHead(node)
	return node.value, true
}

// Adds a Key to the cache if it does not exist, else updates the key's value
// Time O(1)
func(l *LRUCache[K, V]) Put(key K, value V) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// if key exists in the cache 
	node, ok := l.cache[key]
	if ok {
		// update the value of the key 
		node.value = value 
		l.removeNode(node)
		l.addToHead(node)
		return 
	}

	// Key does not exist
	// check the cache capacity 
	if len(l.cache) >= l.capacity {
		// we need to evict LRU node 
		l.evict() 
	}
	
	newNode := &Node[K, V]{key: key, value: value}
	l.cache[key] = newNode
	l.addToHead(newNode) 
}


// Removes a node from the doubly linked list 
func(l *LRUCache[K, V]) removeNode(node *Node[K, V]) {
	// Adjust the link between node's previous node and node's next node. 

	// Defensive nil-checks (although we have head and tail of the list)
	if node.prev != nil {
		node.prev.next = node.next 
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	// Set node's links to nil 
	node.prev = nil 
	node.next = nil 
}

// Adds a node to the head of the doubly linked list 
func(l *LRUCache[K, V]) addToHead(node *Node[K, V]) {
	// fix the link between current Node and the node next to l.Head 
	node.next = l.head.next 
	l.head.next.prev = node 
	// fix the link between current Node and l.Head
	l.head.next = node 
	node.prev = l.head 
}

// Evicts LRU node from the cache 
func(l *LRUCache[K, V]) evict() {
	evictedNode := l.tail.prev 
	// cache is not empty 
	if evictedNode != l.head {
		evictedKey := evictedNode.key
		l.removeNode(evictedNode)
		// Delete this key from cache 
		fmt.Println("Evicted key from cache ", evictedKey)
		delete(l.cache, evictedKey)
	}
}

// Prints Cache from MRU to LRU
func(l *LRUCache[K, V]) DisplayCache() {
	l.mu.Lock()
	defer l.mu.Unlock()

	fmt.Println("Printing LRU cache ")
	for cur := l.head.next ; cur != l.tail ; cur = cur.next {
		fmt.Printf("key: %v, value: %v\n", cur.key, cur.value)
	}
}

func (l *LRUCache[K, V]) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()

	return len(l.cache)
}

func (l *LRUCache[K, V]) KeysMRUToLRU() []K {
	l.mu.Lock()
	defer l.mu.Unlock() 

	keys := make([]K, 0, len(l.cache))

	for cur := l.head.next ; cur != l.tail ; cur = cur.next {
		keys = append(keys, cur.key)
	}
	return keys 
}