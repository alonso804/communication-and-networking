from edge import Edge
from vertex import Vertex
from queue import Queue
import math

def prettyPrint(title, matrix):
    print('\t' + title + ':')

    for i in range(len(matrix)):
        for j in range(len(matrix)):
            print('\t' + str(matrix[i][j]), end='')
        print()

def printByStep(i, distance, path):
    print("Step", str(i) + ':')
    prettyPrint('Distance', distance)
    print()
    prettyPrint('Path', path)


class Graph:
    def __init__(self):
        self.nodesCount = 0
        self.edgesCount = 0
        self.nodes = {}
        self.edges = {}

    def createVertex(self, id):
        if id in self.nodes:
            print("Vertex already exist!")
            return

        vertex = Vertex(id)

        self.nodes[id] = vertex
        self.nodesCount += 1

    def createEdge(self, startId, endId, weight):
        if startId == endId:
            print("Can not point to itself!")
            return

        if startId not in self.nodes or endId not in self.nodes:
            print("One vertex does not exist!")
            return

        if endId in self.nodes[startId].edges:
            print("Edge already exists!")

        vertexStart = self.nodes[startId]
        vertexEnd = self.nodes[endId]

        edge = Edge(self.edgesCount, weight)

        vertexStart.edges[endId] = edge
        vertexEnd.edges[startId] = edge

        self.edges[self.edgesCount] = (startId, endId, weight)
        self.edgesCount += 1

    def floydWarshall(self):
        distance = []
        path = []
        hasher = {}

        index = 0
        for node in self.nodes.keys():
            hasher[index] = node
            index += 1

        for i in range(self.nodesCount):
            tempPath = []
            tempDistance = []

            for j in range(self.nodesCount):
                tempPath.append(hasher[j])
                if i == j:
                    tempDistance.append(0)
                else:
                    currentEdges = self.nodes[hasher[i]].edges
                    currentDistance = math.inf
                    if hasher[j] in currentEdges:
                        currentDistance = currentEdges[hasher[j]].weight

                    tempDistance.append(currentDistance)

            path.append(tempPath)
            distance.append(tempDistance)

        printByStep(1, distance, path)

        iteration = 0
        for k in range(self.nodesCount):
            for i in range(self.nodesCount):
                for j in range(self.nodesCount):
                    if distance[i][k] + distance[k][j] < distance[i][j]:
                        path[i][j] = '\"' + hasher[iteration] + '\"'
                        distance[i][j] = distance[i][k] + distance[k][j]
            iteration += 1
            printByStep(k + 2, distance, path)

        # printByStep(2, distance, path)

    def output(self):
        for k, v in self.nodes.items():
            print(self.nodes[k])
