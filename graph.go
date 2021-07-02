package main

import (
	"fmt"
	"math"
)

type graph struct {
	nodes       map[string]*vertex
	countVertex int
	countEdges  int
}

func newGraph() graph {
	g := graph{
		nodes:       make(map[string]*vertex),
		countVertex: 0,
		countEdges:  0,
	}

	return g
}

func (g *graph) insertVertex(id string, data float64) {
	if _, exist := g.nodes[id]; exist {
		fmt.Printf("Vertex '%s' already exists!\n", id)
		return
	}

	tempVertex := newVertex(data)

	g.nodes[id] = &tempVertex
	g.countVertex++
}

func (g *graph) insertEdge(startId, endId string, data float64) {
	if startId == endId {
		fmt.Println("Can not point to itself!")
		return
	}

	startVertex, startExist := g.nodes[startId]
	endVertex, endExist := g.nodes[endId]

	if !startExist || !endExist {
		fmt.Println("One vertex does not exist!")
		return
	}

	if _, exist := startVertex.edges[endId]; exist {
		fmt.Printf("Edge between '%s' and '%s' already exists!\n", startId, endId)
		return
	}

	tempEdge := newEdge(data)

	startVertex.edges[endId] = &tempEdge
	endVertex.edges[startId] = &tempEdge

	g.countEdges++
}

func (g *graph) deleteEdge(startId, endId string) {
	if startId == endId {
		fmt.Println("Can not be same vertex!")
		return
	}

	startVertex, startExist := g.nodes[startId]
	endVertex, endExist := g.nodes[endId]

	if startExist && endExist {
		_, exist := startVertex.edges[endId]
		if exist {
			delete(startVertex.edges, endId)

			delete(endVertex.edges, startId)

			g.countEdges--
			return
		}

		fmt.Println("Edge does not exist!")
		return
	}

	fmt.Println("One vertex does not exist!")
}

func (g *graph) deleteVertex(id string) {
	if vertexId, exist := g.nodes[id]; exist {

		for edgeId := range vertexId.edges {
			g.deleteEdge(id, edgeId)
		}

		delete(g.nodes, id)

		g.countVertex--
		return
	}

	fmt.Printf("Vertex '%s' does not exist!\n", id)
}

func (g *graph) bellmanFord(idStart string, output bool) graph {
	if _, exist := g.nodes[idStart]; !exist {
		fmt.Println("Vertex does not exist!")
		return newGraph()
	}

	bell := make(map[string]struct {
		string
		float64
	})

	INT_MAX := math.Inf(1)

	for vertexId := range g.nodes {
		bell[vertexId] = struct {
			string
			float64
		}{idStart, INT_MAX}
	}

	bell[idStart] = struct {
		string
		float64
	}{idStart, 0}

	for vertexId, distance := range g.nodes[idStart].edges {
		bell[vertexId] = struct {
			string
			float64
		}{idStart, (*distance).data}
	}

	for i := 0; i < 5; i++ {
		for vertexId, vertexValue := range g.nodes {
			if bell[vertexId].float64 != INT_MAX {
				for edgeId, edgeDistance := range (*vertexValue).edges {
					if bell[vertexId].float64+(*edgeDistance).data <= bell[edgeId].float64 {
						bell[edgeId] = struct {
							string
							float64
						}{vertexId, bell[vertexId].float64 + (*edgeDistance).data}
					}
				}
			}
		}
	}

	result := newGraph()

	for vertexId, vertexValue := range g.nodes {
		result.insertVertex(vertexId, vertexValue.data)
	}

	for to, from := range bell {
		if to != from.string && from.float64 != INT_MAX {
			result.insertEdge(to, from.string, (*g.nodes[from.string]).edges[to].data)
		}
	}

	if output {
		for to, from := range bell {
			if from.float64 == INT_MAX {
				fmt.Printf("'%s' and '%s' are not connected!\n", to, from.string)
			} else {
				fmt.Printf("'%s' coming from '%s' with %.2f\n", to, from.string, from.float64)
			}
		}
	}

	return result
}

func (g *graph) print() {
	for vertexId, vertexValue := range g.nodes {
		fmt.Printf("%s\n", vertexId)
		for edgeId, edgeDistance := range (*vertexValue).edges {
			fmt.Printf("  • %s: %.2f\n", edgeId, edgeDistance.data)
		}
	}
}
