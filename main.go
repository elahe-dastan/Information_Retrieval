package main

import (
	"Information_Retrieval/index"
	vector_space "Information_Retrieval/vector-space"
)

func main() {
	i := index.NewIndex("./docs", 6)
	indexFile := i.Construct()
	v := vector_space.NewVectorizer(indexFile, 3)
	v.Vectorize()
	v.Query("نشست کمیسیون")
	//fmt.Println(query.Query())
}
