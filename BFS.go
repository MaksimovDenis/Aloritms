package main

import "fmt"

type Queue struct {
	List []int
}

// Функция добавления элемента в очередь
func (q *Queue) Enqueue(element int) {
	q.List = append(q.List, element)
}

// Функция удаления элемента из очереди
func (q *Queue) Dequeue() int {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return 0
	}
	element := q.List[0]
	q.List = q.List[1:]
	return element
}

// Функция проверки очереди на наличие в ней элементов
func (q *Queue) isEmpty() bool {
	return len(q.List) == 0
}

// BFS() через матрицу смежности
func BFSmatrix(graph [][]int, node int) {
	//Иницилизируем мапу, которая содержит информация
	//Была ли посещен та или иная вершина
	isVisisted := make(map[int]bool)
	//Создаём очередь, в которую будут добавлены элемнеты, которые
	//находятся на одном и том же уровен, что и вершина
	var bfsQueue Queue

	//Помечаем, что данную вершину мы посетили
	isVisisted[node] = true

	//Добавляем вершину в очередь
	bfsQueue.Enqueue(node)

	//Запускаем цикл, пока очередь не закончится
	for !bfsQueue.isEmpty() {
		currNode := bfsQueue.List[0]
		fmt.Print(currNode, " ")
		//Добавляем вершины, которые не посетили в матрицу
		for i := 0; i < len(graph[currNode]); i++ {
			if graph[currNode][i] == 1 && !isVisisted[i] {
				bfsQueue.Enqueue(i)
				isVisisted[i] = true
			}
		}
		//Удаляем вершину, которую рассматривали
		bfsQueue.Dequeue()
	}

}

// BFS() через матрицу смежных вершин
func BFSlist(graph [4][]int, node int) {
	//Иницилизируем мапу, которая содержит информация
	//Была ли посещен та или иная вершина
	isVisisted := make(map[int]bool)
	//Создаём очередь, в которую будут добавлены элемнеты, которые
	//находятся на одном и том же уровен, что и вершина
	var bfsQueue Queue

	//Помечаем, что данную вершину мы посетили
	isVisisted[node] = true

	//Добавляем вершину в очередь
	bfsQueue.Enqueue(node)

	//Запускаем цикл, пока очередь не закончится
	for !bfsQueue.isEmpty() {
		currNode := bfsQueue.List[0]
		fmt.Print(currNode, " ")
		//Добавляем вершины, которые не посетили в матрицу
		for _, nodes := range graph[currNode] {
			if !isVisisted[nodes] {
				bfsQueue.Enqueue(nodes)
				isVisisted[nodes] = true
			}
		}
		//Удаляем вершину, которую рассматривали
		bfsQueue.Dequeue()
	}

}
