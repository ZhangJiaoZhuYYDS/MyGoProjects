package main

func main() {

}

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}
func (this *MyStack) Pop() int {
	n := len(this.queue) - 1 // 判断长度
	for n != 0 {             // 除了最后一个，其他的全部删除再重新添加到队列末尾
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, val)
		n--
	}
	// 弹出元素
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}
func (this *MyStack) Top() int {
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
