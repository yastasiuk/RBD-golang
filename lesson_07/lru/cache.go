package lru

type LruCache interface {
	// Якщо наш кеш вже повний (ми досягли нашого capacity)
	// то має видалитись той елемент, який ми до якого ми доступались (читали) найдавніше
	Put(key, value string)
	Get(key string) (string, bool)
}

type cacheValue struct {
	value string
	node  *Node
}

type Node struct {
	key      string
	previous *Node
	next     *Node
}

type cache struct {
	capacity    int
	currentSize int
	positionMap map[string]cacheValue
	first       *Node
	last        *Node
}

func (c *cache) Put(key, value string) {
	// Cache already exists - do nothing
	if _, ok := c.Get(key); ok {
		c.positionMap[key] = cacheValue{value, c.positionMap[key].node}
	}

	c.currentSize += 1

	newNode := &Node{
		key:      key,
		previous: nil,
		next:     nil,
	}
	newCacheValue := cacheValue{
		value: value,
		node:  newNode,
	}
	c.positionMap[key] = newCacheValue

	// Repositioning last nodes
	if c.first == nil {
		c.first = newNode
	}
	lastNode := c.last
	if lastNode != nil {
		lastNode.next = newNode
		newNode.previous = lastNode
	}
	c.last = newNode

	// Cleaning up old cache values
	if c.currentSize > c.capacity {
		firstNode := c.first
		c.first = firstNode.next
		if c.first != nil {
			c.first.previous = nil
		}
		delete(c.positionMap, firstNode.key)
		c.currentSize -= 1
	}
}

func (c *cache) Get(key string) (string, bool) {
	cachedValue, ok := c.positionMap[key]
	if !ok {
		return "", false
	}

	currentNode := cachedValue.node

	previousNode := currentNode.previous
	nextNode := currentNode.next

	// Merging prev & next nodes
	if previousNode != nil {
		previousNode.next = nextNode
	}
	if nextNode != nil {
		nextNode.previous = previousNode
	}

	// Moving current node to last position

	if c.first == currentNode && currentNode.next != nil {
		c.first = currentNode.next
	}
	lastNode := c.last
	lastNode.next = currentNode
	currentNode.previous = lastNode
	c.last = currentNode
	currentNode.next = nil

	return cachedValue.value, true
}

func NewLruCache(capacity int) LruCache {
	var cacheImplementation LruCache

	cacheImplementation = &cache{
		capacity:    capacity,
		currentSize: 0,
		positionMap: make(map[string]cacheValue),
		first:       nil,
		last:        nil,
	}

	return cacheImplementation
}
