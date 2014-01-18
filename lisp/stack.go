package lisp

type Stack struct{
	top interface{}
	size int
	items []interface{}
}

func New() *Stack{
	return &Stack{
		nil,
		0,
		nil,
	}
}

func (stack *Stack) Push(v interface{}){
	stack.items = append(stack.items, v)
	stack.size += 1
}

func (stack *Stack) Pop() interface{}{
	top := stack.Top()
	stack.items[stack.Size() - 1] = nil
	stack.size -= 1
	return top
}

func (stack *Stack) Top() interface{}{
	return stack.items[stack.Size() - 1]
}

func (stack *Stack) Empty() bool{
	return stack.size == 0
}

func (stack *Stack) Size() int{
	if stack == nil{
		return 0
	}
	return stack.size
}

