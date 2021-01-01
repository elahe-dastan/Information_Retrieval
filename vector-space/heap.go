package vector_space

type Similarity struct {
	DocId int
	Cos   float64
}

// A Heap is a max-heap of Similarity.
type Heap []Similarity

func (h Heap) Len() int           { return len(h) }
// changed the less function so it became a max heap
func (h Heap) Less(i, j int) bool { return h[i].Cos > h[j].Cos }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Similarity))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
