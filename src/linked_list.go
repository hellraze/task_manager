package src

import "errors"

var (
	ErrorElementIsNotFound   = errors.New("element not found")
	ErrorThisIsFirstElement  = errors.New("this is first element")
	ErrorPrevElementNotFound = errors.New("previous element not found")
)

type Node struct {
	id          int
	priority    int
	time        int
	description string
	next        *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Insert(node *Node) {

	if l.Head == nil || node.priority > l.Head.priority {
		node.next = l.Head
		l.Head = node
	} else {
		current := l.Head

		for current.next != nil && node.priority <= current.next.priority {
			current = current.next
		}

		node.next = current.next
		current.next = node
	}
}

func (l *LinkedList) Delete(id int) {
	prevNode, err := l.FindPrevNodeById(id)
	if errors.Is(err, ErrorThisIsFirstElement) {
		l.Head = l.Head.next
		return
	}
	prevNode.next = prevNode.next.next

}

func (l *LinkedList) FindPrevNodeById(id int) (*Node, error) {
	current := l.Head
	if current.id == id {
		return nil, ErrorThisIsFirstElement
	}
	for current != nil {
		if current.next.id == id {
			return current, nil
		}
		current = current.next
	}
	return nil, ErrorPrevElementNotFound
}

func (l *LinkedList) FindById(id int) (*Node, error) {
	current := l.Head
	for current != nil {
		if current.id == id {
			return current, nil
		}
		current = current.next
	}
	return nil, ErrorElementIsNotFound
}
