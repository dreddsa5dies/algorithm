package queue

type (
	// Queue -  Очередь представляется в качестве линейного списка,
	// в котором добавление/удаление элементов идет строго с соответствующих его концов.
	Queue struct {
		start, end *node
		length     int
	}
	node struct {
		value interface{}
		next  *node
	}
)

// New - создание новой очереди
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Dequeue - Удаление элемента из передней части очереди и возврат его значения.
func (que *Queue) Dequeue() interface{} {
	if que.length == 0 {
		return nil
	}
	n := que.start
	if que.length == 1 {
		que.start = nil
		que.end = nil
	} else {
		que.start = que.start.next
	}
	que.length--
	return n.value
}

// Enqueue - Добавить новый элемент в конец очереди.
func (que *Queue) Enqueue(value interface{}) {
	n := &node{value, nil}
	if que.length == 0 {
		que.start = n
		que.end = n
	} else {
		que.end.next = n
		que.end = n
	}
	que.length++
}

// Len - Возвращает количество элементов внутри очереди.
func (que *Queue) Len() int {
	return que.length
}

// Peek - Вернуть значение элемента в начале очереди, не удаляя его
func (que *Queue) Peek() interface{} {
	if que.length == 0 {
		return nil
	}
	return que.start.value
}
