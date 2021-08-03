from graph import Graph
from edge import Edge
from queue import Queue

if __name__ == "__main__":
    graph = Graph()

    graph.createVertex("A")
    graph.createVertex("B")
    graph.createVertex("C")
    graph.createVertex("D")
    graph.createVertex("E")

    graph.createEdge("A", "B", 2)
    graph.createEdge("A", "C", 1)
    graph.createEdge("A", "D", 4)

    graph.createEdge("B", "D", 3)
    graph.createEdge("B", "E", 4)

    graph.createEdge("C", "D", 2)

    graph.createEdge("D", "E", 2)

    graph.output()
    print()
    graph.floydWarshall()
