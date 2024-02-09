package main

import (
	"fmt"
	"math"
)

// Алгоритм Дейкстры
type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	node     int
	distance float64
	edges    []*Edges
	visited  bool
}

type Edges struct {
	weight float64
	target *Vertex
}

// Создаём вершину
func (g *Graph) AddVertex(node int) {
	g.vertices = append(g.vertices, &Vertex{node: node})
}

// Создаём ребро
func (g *Graph) AddEdge(sourceNode, targetNode int, weight float64) {
	source := g.FindVertex(sourceNode) //Исходящая
	target := g.FindVertex(targetNode) //Смежная вершина
	source.edges = append(source.edges, &Edges{weight: weight, target: target})
}

// Ищем необходмимую нам вершину
func (g *Graph) FindVertex(node int) *Vertex {
	for _, v := range g.vertices {
		if v.node == node {
			return v
		}
	}
	return nil
}

func (g *Graph) Dijkstra(sourceNode int) {
	source := g.FindVertex(sourceNode)
	if source == nil {
		fmt.Println("Исходной вершины не существует")
		return
	}

	//Проходися по структуре и в случае аналичия отрицательных значений меняем их на положительные
	for _, v := range g.vertices {
		v.distance = math.Inf(1)
		v.visited = false
	}

	source.distance = 0

	for _, v := range g.vertices {
		v = g.minDistanceVertex()
		v.visited = true
		for _, e := range v.edges {
			if !e.target.visited {
				newDist := v.distance + e.weight
				if newDist < e.target.distance {
					e.target.distance = newDist
				}
			}
		}
	}
}

func (g *Graph) minDistanceVertex() *Vertex {
	var minVertex *Vertex
	minDist := math.Inf(1)

	for _, v := range g.vertices {
		if !v.visited && v.distance < minDist {
			minVertex = v
			minDist = v.distance
		}
	}
	return minVertex
}
