/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    pointers [][]*NestedInteger
    indexes  []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{
        pointers: [][]*NestedInteger{nestedList},
        indexes:  []int{-1},
    }
}

func (this *NestedIterator) Next() int {
    level := len(this.pointers)-1
    return this.pointers[level][this.indexes[level]].GetInteger()
}

func (this *NestedIterator) HasNext() bool {
    level := len(this.pointers) - 1
    if level < 0 {
        return false
    }
    if this.indexes[level] + 1 >= len(this.pointers[level]) {
        this.pointers = this.pointers[:level]
        this.indexes = this.indexes[:level]
        return this.HasNext()
    }
    this.indexes[level]++
    if this.pointers[level][this.indexes[level]].IsInteger() {
        return true
    }
    this.pointers = append(this.pointers, this.pointers[level][this.indexes[level]].GetList())
    this.indexes = append(this.indexes, -1)
    return this.HasNext()
}