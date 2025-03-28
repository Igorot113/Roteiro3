package main

/*adjiasjdoa*/
import (
	exerciciotodolist "Roteiro3/ExercicioToDoList"
	roteiro3_1 "Roteiro3/Roteiro3_1"
	"fmt"
)

func main() {
	var b float64 = 10.50
	var p2 *float64 = &b
	fmt.Println("\nValor de b antes da modificação:", b)
	*p2 = 20.50
	fmt.Println("Valor de b antes da modificação:", b)
	var a int = 10
	fmt.Println("Fora da função increment:", a)
	increment(a)
	var c int = 20
	fmt.Println("Fora da função incrementByPointer:", c)
	incrementByPointer(&c)
	x1, x2 := 10, 20
	fmt.Println("Antes do swap:")
	fmt.Print("x = ", x1)
	fmt.Print("\n")
	fmt.Print("y = ", x2)
	swap(x1, x2)
	list := LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)

	fmt.Println("\nLista Encadeada Simples:")
	list.Print()

	list.RemoveFirst()
	fmt.Println("Após remover o primeiro:")
	list.Print()

	listDC := roteiro3_1.DoublyLinkedList{}
	listDC.Append(1)
	listDC.Append(2)
	listDC.Append(3)

	fmt.Println("Lista Duplamente Encadeada do inicio para o fim:")
	listDC.PrintForward()

	fmt.Println("Lista Duplamente Encadeada do fim para o inicio:")
	listDC.PrintBackward()

	listDC.RemoveLast()
	fmt.Println("\nApos remover o ultimo elemento:")
	listDC.PrintForward()

	listExercicio := exerciciotodolist.DoublyLinkedList2{}
	listExercicio.Append("Estudar GO")
	listExercicio.Append("Fazer compras")
	listExercicio.Append("Academia")

	fmt.Println("\nExercicio de Lista Encadeada: do inicio para o fim:")
	listExercicio.PrintForward()

	fmt.Println("\nExercicio de Lista Encadeada: do fim para o inicio:")
	listExercicio.PrintBackward()

	fmt.Println("\nApos remover o ultimo elemento:")
	listExercicio.RemoveLast()

	fmt.Println("\nExercicio de Lista Encadeada: do inicio para o fim:")
	listExercicio.PrintForward()

	fmt.Println("\nExercicio de Lista Encadeada: do fim para o inicio:")
	listExercicio.PrintBackward()

}

func increment(val int) {
	val++
	fmt.Println("Dentro da função increment:", val)
}

func incrementByPointer(val *int) {
	*val++
	fmt.Println("Dentro da função incrementByPointer:", *val)
}

func swap(x1, x2 int) {
	var val1 *int = &x1
	var val2 *int = &x2
	val1, val2 = &x2, &x1
	fmt.Println("\nDepois do swap: ")
	fmt.Print("x = ", *val1)
	fmt.Print("\n")
	fmt.Print("y = ", *val2)
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) RemoveFirst() {
	if l.head != nil {
		l.head = l.head.next
	}
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, "->")
		current = current.next
	}
	fmt.Println("nil")
}

/*O que esse código faz?
1. Criamos uma estrutura Node para armazenar os dados.
2.
*/
