package exerciciotodolist

import "fmt"

type Task struct {
	data string
	next *Task
	prev *Task
}

type DoublyLinkedList2 struct {
	head *Task
	tail *Task
}

func (l *DoublyLinkedList2) Append(data string) {
	newTask := &Task{data: data}
	if l.tail == nil {
		l.head = newTask
		l.tail = newTask
		return
	}
	l.tail.next = newTask
	newTask.prev = l.tail
	l.tail = newTask
}

func (l *DoublyLinkedList2) RemoveLast() {
	if l.tail == nil {
		return
	}
	if l.tail == l.head {
		l.head = nil
		l.tail = nil
		return
	}
	l.tail = l.tail.prev
	l.tail.next = nil
}

func (l *DoublyLinkedList2) PrintForward() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " <-> ")
		current = current.next
	}
	fmt.Println(" nil")
}

func (l *DoublyLinkedList2) PrintBackward() {
	current := l.tail
	for current != nil {
		fmt.Print(current.data, " <-> ")
		current = current.prev
	}
	fmt.Println(" nil")
}
