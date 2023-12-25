package main

func main() {

}

type MyQueue struct {
	stackIn  []int // 输入栈
	stackOut []int // 输出栈
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

// push 入栈
func (this *MyQueue) Push(i int) {
	this.stackIn = append(this.stackIn, i)
}

// pop 出栈
// 在输出栈执行出栈操作时，判断输出栈是否为空；若为空，就把输入栈数据全部导入输出栈；若输入栈也为空，就返回-1，代表空栈
func (this *MyQueue) Pop() int {
	// 获取输入栈和输出栈的长度
	inLen, outLen := len(this.stackIn), len(this.stackOut)
	if outLen == 0 {
		if inLen == 0 {
			return -1
		}
		//倒叙遍历输入栈，把元素添加到输出栈
		for i := inLen - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		//清空输入栈，重新计算长度
		this.stackIn = []int{}
		inLen = len(this.stackIn)
		outLen = len(this.stackOut)
	}
	//出栈
	val := this.stackOut[outLen-1]
	this.stackOut = this.stackOut[:outLen-1]
	return val
}

// 获取队列第一个元素：先调用弹栈，获取弹栈元素，然后添加弹栈元素，返会弹栈元素
func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == -1 {
		return -1
	}
	this.stackOut = append(this.stackOut, val)
	return val
}

// 是否为空
func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}
