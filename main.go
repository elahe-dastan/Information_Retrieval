package main

import (
	vector_space "Information_Retrieval/vector-space"
	"container/heap"
	"fmt"
)

func main() {
	s := &vector_space.Heap{vector_space.Similarity{
		DocId: 1,
		Cos:   21,
	},vector_space.Similarity{
		DocId: 4,
		Cos:   34,
	},vector_space.Similarity{
		DocId: 5,
		Cos:   63,
	},
	vector_space.Similarity{
		DocId: 3,
		Cos:   65,
	}}

	heap.Init(s)
	heap.Push(s, vector_space.Similarity{
		DocId: 7,
		Cos:   33,
	})

	for s.Len() > 0 {
		fmt.Println(heap.Pop(s))
	}
	//i := index.NewIndex("./docs", 6)
	//indexFile := i.Construct()
	//v := vector_space.NewVectorizer(indexFile, 3)
	//v.Vectorize()
	//fmt.Println(query.Query("نشست کمیسیون"))
}
