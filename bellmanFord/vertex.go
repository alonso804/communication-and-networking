package main

type vertex struct {
	edges map[string]*edge
	data  float64
}

func newVertex(data float64) vertex {
	v := vertex{
		edges: make(map[string]*edge),
		data:  data,
	}

	return v
}
