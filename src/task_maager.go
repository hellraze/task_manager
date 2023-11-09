package src

import (
	"fmt"
	"strings"
)

type TaskManager struct {
	Queue *PriorityQueue
}

func (t *TaskManager) SelectTasksByTimer(freeTime int) string {
	var (
		tasks   []string
		counter int
	)
	current := t.Queue.List.Head
	for current != nil {
		if freeTime >= current.time {
			freeTime -= current.time
			tasks = append(tasks, current.description)
		}
		current = current.next
		counter += 1
	}
	result := strings.Join(tasks, ", ")
	if counter == 0 {
		return "К сожалению, вы не успеете решить ни одной задачи"
	}
	return result
}

func (t *TaskManager) View() {
	x := t.Queue.List.Head
	for x.next != nil {
		fmt.Printf("id: %d Описание: %v, Приоритет: %d, Это займет: %d минут\n", x.id, x.description, x.priority, x.time)
		x = x.next
	}
	fmt.Printf("id: %d Описание: %v, Приоритет: %d, Это займет: %d минут\n", x.id, x.description, x.priority, x.time)
}

func (t *TaskManager) Insert(x string, priority int, time int) {
	t.Queue.Enqueue(x, priority, time)
}

func (t *TaskManager) Delete() string {
	temp := t.Queue.Dequeue()
	return fmt.Sprintf("Выполнено: %v за %d минут", temp.description, temp.time)
}

func NewTaskManager() *TaskManager {
	linkedList := &LinkedList{}
	priorityQueue := &PriorityQueue{
		List: linkedList,
	}

	return &TaskManager{
		Queue: priorityQueue,
	}
}
