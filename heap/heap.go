package heap

type Similarity struct {
	DocId int
	Cos   float64
}

// A SimilarityHeap is a max-heap of Similarity.
type SimilarityHeap []Similarity

func (h SimilarityHeap) Len() int { return len(h) }
// changed the less function so it became a max heap
func (h SimilarityHeap) Less(i, j int) bool { return h[i].Cos > h[j].Cos }
func (h SimilarityHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SimilarityHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Similarity))
}

func (h *SimilarityHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Frequency struct {
	DocId string
	Freq  int
}

// A FrequencyHeap is a max-heap of Frequency
type FrequencyHeap []Frequency

func (f FrequencyHeap) Len() int { return len(f) }
// changed the less function so it became a max heap
func (f FrequencyHeap) Less(i, j int) bool { return f[i].Freq > f[j].Freq }
func (f FrequencyHeap) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func (f *FrequencyHeap) Pop() interface{} {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}

func (f *FrequencyHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*f = append(*f, x.(Frequency))
}
