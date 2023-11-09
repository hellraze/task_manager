package src

import "errors"

var lastId int

type PriorityQueue struct {
	List *LinkedList
}

func (q *PriorityQueue) Enqueue(x string, priority int, time int) {
	lastId++
	newNode := &Node{id: lastId, description: x, priority: priority, time: time}

	q.List.Insert(newNode)
}

func (q *PriorityQueue) Dequeue() *Node {
	temp := q.List.Head
	q.List.Head = temp.next

	return temp
}

func (q *PriorityQueue) UpdatePriority(id int, priority int) {
	current, err := q.List.FindById(id)
	if errors.Is(err, ErrorElementIsNotFound) {
		return
	}
	current.priority = priority
	q.List.Delete(id)
	q.List.Insert(current)
}
