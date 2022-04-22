type Node struct {
    key int
    value int
    next *Node
}

type MyHashMap struct {
    items []*Node
    size int
}

func Constructor() MyHashMap {
    size := 1000
    return MyHashMap {
        make([]*Node, size),
        size,
    }
}

func (this *MyHashMap) Put(key int, value int)  {
    index := key % this.size
    item := this.items[index]
    
    if item == nil {
        this.items[index] = &Node{key: key, value: value}
        return
    }
    
    for item != nil {
        if item.key == key {
            item.value = value
            return
        }
        
        if item.next == nil {
            item.next = &Node{key: key, value: value}
            return
        }
        
        item = item.next
    }
}

func (this *MyHashMap) Get(key int) int {
    index := key % this.size
    if this.items[index] == nil {
        return -1
    }
    
    item := this.items[index]
    
    for item != nil {
        if item.key == key {
            return item.value
        }
        
        item = item.next
    }
    
    return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    index := key % this.size
    if this.items[index] == nil {
        return
    }
    
    item := this.items[index]
    
    if item.key == key {
        this.items[index] = item.next
        return
    }
    
    prev, item := item, item.next
    
    for item != nil {
        if item.key == key {
            prev.next = item.next
            return
        }
        
        prev = item        
        item = item.next
    }
}
