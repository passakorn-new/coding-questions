type MyHashMap struct {
    Size int
    Hashmap [1e6][2]int
}

func Constructor() MyHashMap {
    var hashmap MyHashMap
    hashmap.Size = 1e6
    hashmap.Hashmap = [1e6][2]int{}
    return hashmap
}

func (this *MyHashMap) Put(key int, value int)  {
    hash := key % this.Size
    this.Hashmap[hash] = [2]int{key, value}
}

func (this *MyHashMap) Get(key int) int {
    hash := key % this.Size
    if this.Hashmap[hash] == [2]int{} {
        return -1
    } 
    return this.Hashmap[hash][1]
}

func (this *MyHashMap) Remove(key int)  {
    hash := key % this.Size
    this.Hashmap[hash] = [2]int{}
}