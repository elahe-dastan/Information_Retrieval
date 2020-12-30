package main

import (
	"Information_Retrieval/index"
	"Information_Retrieval/query"
	"fmt"
)

func main() {
	i := index.NewIndex("./docs", 6)
	indexFile := i.Construct()


	fmt.Println(query.Query("نشست کمیسیون"))
}
