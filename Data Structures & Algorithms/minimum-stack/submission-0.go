type MinStack struct {
	top *Node
	min *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.top == nil {
		this.top = &Node{
			val: val,
		}
		this.min = &Node{
			val: val,
		}
		return
	}

	this.top = &Node{
		val: val,
		prev: this.top,
	}
	this.min = &Node{
		val: min(val, this.min.val),
		prev: this.min,
	}
}

func (this *MinStack) Pop() {
	this.top = this.top.prev
	this.min = this.min.prev
}

func (this *MinStack) Top() int {
	return this.top.val
}

func (this *MinStack) GetMin() int {
	return this.min.val
}

type Node struct {
	prev *Node
	val int
}