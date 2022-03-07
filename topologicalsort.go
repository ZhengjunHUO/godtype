package godtype

type Dag struct {
	// the size of the graph
	Size int
	// adjacent table for each vertex
	AdjTable map[int][]int
	// stock the result of sort, "source" vertex on the top
	// the vertex "pointed to" at the bottom
	Sorted *Stack
	// to indicate if the vertex as been visited
	visited []bool
}

// initialize the graph
func NewDag(size int) *Dag {
	return &Dag{
		Size: size,
		AdjTable: make(map[int][]int),
		Sorted: NewStack(),
		visited: make([]bool, size),
	}
}

// build the adjacent table
func (d *Dag) AddEdge(fromV, toV int) {
	if _, ok := d.AdjTable[fromV]; !ok {
		d.AdjTable[fromV] = []int{toV}
	}else{
		d.AdjTable[fromV] = append(d.AdjTable[fromV], toV)
	}
}

// recursively go through the current vertex' neighbours, and save the result to stack
func (d *Dag) findPath(vertex int) {
	d.visited[vertex] = true

	if v, ok := d.AdjTable[vertex]; ok {
		for i := range v {
			if !d.visited[v[i]] {
				d.findPath(v[i])
			}
		}
	}

	d.Sorted.Push(vertex)
}

func (d *Dag) TopologicalSort() {
	for i:=0; i<d.Size; i++ {
		if !d.visited[i] {
			d.findPath(i)
		}
	}

}
