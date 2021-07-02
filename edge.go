package main

type edge struct {
	data float64
}

func newEdge(data float64) edge {
	e := edge{
		data: data,
	}

	return e
}
