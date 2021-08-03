import math

class Vertex:
    def __init__(self, id):
        self.id = id
        self.edges = {}

    def __str__(self):
        edges = { key : value.weight for (key,value) in self.edges.items() }

        vertex_content = f'{self.id} -> {edges}'
        return vertex_content
