package stack

type (
	// Stack - список элементов
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// New - Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Len - Return the number of items in the stack
func (st *Stack) Len() int {
	return st.length
}

// Peek - View the top item on the stack
func (st *Stack) Peek() interface{} {
	if st.length == 0 {
		return nil
	}
	return st.top.value
}

// Pop - the top item of the stack and return it
func (st *Stack) Pop() interface{} {
	if st.length == 0 {
		return nil
	}

	n := st.top
	st.top = n.prev
	st.length--
	return n.value
}

// Push - a value onto the top of the stack
func (st *Stack) Push(value interface{}) {
	n := &node{value, st.top}
	st.top = n
	st.length++
}
