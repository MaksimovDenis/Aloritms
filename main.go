package main

import (
	"fmt"
)

func main() {

	fmt.Println("Бинарное дерево")
	root := BinaryTree{}
	root.Insert(5)
	root.Insert(7)
	root.Insert(8)
	root.Insert(1)
	root.Insert(0)
	root.Insert(3)
	root.Insert(2)

	root.PrintInorderMethod()
	fmt.Println()
	fmt.Println("После удаления элемента 5")
	root.Delete(5)
	root.PrintInorderMethod()
	fmt.Println()
	fmt.Println("Поиск в ширину через матрицу смежности")
	graph_matrix := [][]int{{0, 1, 0, 1},
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{1, 0, 1, 0}}

	BFSmatrix(graph_matrix, 2)

	fmt.Println()
	fmt.Println("Поиск в ширину через матрицу смежных вершин")
	var graph_list [4][]int
	graph_list[0] = []int{1, 3}
	graph_list[1] = []int{0, 2}
	graph_list[2] = []int{1, 3}
	graph_list[3] = []int{0, 2}

	BFSlist(graph_list, 2)

	fmt.Println()
	fmt.Println("Алгоритм Дейкстры")
	g := Graph{}

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 1)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 2)
	g.AddEdge(2, 4, 4)
	g.AddEdge(3, 4, 1)

	g.Dijkstra(0)

	for _, v := range g.vertices {
		fmt.Printf("Кратчайшее расстаяние из 0 в %d: %.0f\n", v.node, v.distance)
	}
}
