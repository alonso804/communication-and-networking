package main

func midterm() {
	g1 := newGraph()
	g1.insertVertex("A", 1)
	g1.insertVertex("B", 2)
	g1.insertVertex("C", 3)
	g1.insertVertex("D", 4)
	g1.insertVertex("E", 5)
	g1.insertVertex("F", 6)
	g1.insertVertex("G", 7)

	g1.insertEdge("A", "B", 2)
	g1.insertEdge("A", "C", 1)
	g1.insertEdge("A", "E", 0.5)

	g1.insertEdge("B", "C", 0.5)
	g1.insertEdge("B", "D", 2)
	g1.insertEdge("B", "E", 3)

	g1.insertEdge("C", "G", 1)

	g1.insertEdge("D", "F", 2)
	g1.insertEdge("D", "G", 3)

	g1.insertEdge("E", "F", 2)

	g1.insertEdge("F", "G", 1)

	//g1.deleteEdge("F", "G")

	bellmanFordGraph1 := g1.bellmanFord("A", false)
	bellmanFordGraph1.print()
	//g1.bellmanFord("A", true)
}

func e1() {
	g1 := newGraph()
	g1.insertVertex("A", 1)
	g1.insertVertex("B", 2)
	g1.insertVertex("C", 3)
	g1.insertVertex("D", 4)
	g1.insertVertex("E", 5)
	g1.insertVertex("F", 6)
	g1.insertVertex("G", 7)
	g1.insertVertex("H", 8)
	g1.insertVertex("I", 9)
	g1.insertVertex("J", 10)

	g1.insertEdge("A", "B", 2)
	g1.insertEdge("A", "C", 3)
	g1.insertEdge("A", "D", 5)

	g1.insertEdge("B", "D", 1.5)
	g1.insertEdge("B", "E", 4)
	g1.insertEdge("B", "F", 1)

	g1.insertEdge("C", "D", 1)
	g1.insertEdge("C", "G", 0.5)

	g1.insertEdge("D", "F", 0.3)
	g1.insertEdge("D", "G", 1)

	g1.insertEdge("E", "F", 2)
	g1.insertEdge("E", "H", 2)
	g1.insertEdge("E", "J", 5)

	g1.insertEdge("F", "G", 2.1)
	g1.insertEdge("F", "H", 3)
	g1.insertEdge("F", "I", 1.7)

	g1.insertEdge("G", "I", 4.2)

	g1.insertEdge("H", "I", 3.5)
	g1.insertEdge("H", "J", 4.6)

	g1.insertEdge("I", "J", 0.2)

	bellmanFordGraph1 := g1.bellmanFord("A", true)
	bellmanFordGraph1.print()

	//g1.bellmanFord("A", true)
}

func main() {
	//g1 := newGraph()

	// Insert Vertex
	//fmt.Println("========INSERT VERTEX===========")
	//g1.insertVertex("A", 1)
	//g1.insertVertex("B", 2)
	//g1.insertVertex("C", 3)
	//fmt.Println(g1)

	// Insert Edge
	//fmt.Println("========INSERT EDGE===========")
	//g1.insertEdge("A", "C", 11)
	//g1.insertEdge("B", "C", 3)
	//g1.insertEdge("B", "A", 5)
	//fmt.Println("========BellmanFord===========")
	//g1.bellmanFord("A")
	//g1.insertEdge("C", "A", 3)
	//fmt.Println(*(g1.nodes["A"]).edges["C"])
	//fmt.Println(*(g1.nodes["A"]).edges["C"].vertexes[0])
	//fmt.Println(*(g1.nodes["A"]).edges["C"].vertexes[1])
	//fmt.Println((g1.nodes["A"]).edges)

	// Delete Edge
	//fmt.Println("========DELETE EDGE===========")
	//g1.deleteEdge("C", "A")
	//fmt.Println((g1.nodes["A"]).edges)

	// Delete Vertex
	//fmt.Println("========DELETE VERTEX===========")
	//g1.deleteVertex("A")
	//fmt.Println(g1)

	//midterm()
	e1()
}
