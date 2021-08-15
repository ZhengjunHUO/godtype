package godtype

type UF struct {
	count	int
	parent	[]int
	size	[]int
}

func NewUF(n int) *UF {
	parent, size := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	return &UF{n, parent, size}
}

func (uf *UF) FindRoot(a int) int {
	for a != uf.parent[a] {
		uf.parent[a] = uf.parent[uf.parent[a]]
		a = uf.parent[a]
	}
	
	return a
}


func (uf *UF) Union(a, b int) {
	ar, br := uf.FindRoot(a), uf.FindRoot(b)
	if ar == br {
		return
	}

	if uf.size[ar] < uf.size[br] {
		uf.parent[ar] = br
		uf.size[br] += uf.size[ar]
	}else{
		uf.parent[br] = ar
		uf.size[ar] += uf.size[br]
	}
	
	uf.count--
}



func (uf *UF) IsLinked(a, b int) bool {
	return uf.FindRoot(a) == uf.FindRoot(b)
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) Parent() []int {
	p := make([]int, len(uf.parent))
	copy(p, uf.parent)
	return p
}

func (uf *UF) Size() []int {
	s := make([]int, len(uf.size))
	copy(s, uf.size)
	return s
}

func (uf *UF) SetCount(c int) {
	uf.count = c
}

func (uf *UF) SetParent(p []int) {
	copy(uf.parent, p)
}

func (uf *UF) SetSize(s []int) {
	copy(uf.size, s)
}
