type MyStack struct {
    queue []int
}


func Constructor() MyStack {
  return MyStack{ []int{} }
}


func (this *MyStack) Push(x int)  {
    this.queue = append([]int{x}, this.queue[0:]...)
}


func (this *MyStack) Pop() int {
    temp := this.queue[0]
    this.queue = this.queue[1:]
    return temp
}


func (this *MyStack) Top() int {
    return this.queue[0] 
}


func (this *MyStack) Empty() bool {
    return len(this.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */