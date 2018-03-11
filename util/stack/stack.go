package stack

type (
	// Stack - список элементов
	// Зачастую стек реализуется в виде однонаправленного списка
	// (каждый элемент в списке содержит помимо хранимой информации в стеке указатель
	// на следующий элемент стека).
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// New - создание нового стэка
func New() *Stack {
	return &Stack{nil, 0}
}

// Len - возврат количества элементов в стеке
func (st *Stack) Len() int {
	return st.length
}

// Peek - возврат верхнего элемента
func (st *Stack) Peek() interface{} {
	if st.length == 0 {
		return nil
	}
	return st.top.value
}

// Pop - возврат элемента и удаление его
func (st *Stack) Pop() interface{} {
	if st.length == 0 {
		return nil
	}

	n := st.top
	st.top = n.prev
	st.length--
	return n.value
}

// Push - значение в верхней части стека
func (st *Stack) Push(value interface{}) {
	n := &node{value, st.top}
	st.top = n
	st.length++
}
