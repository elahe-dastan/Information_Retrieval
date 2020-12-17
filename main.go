package main

import "Information_Retrieval/index"

func main() {
	i := index.NewIndex("./docs", 6)
	i.Construct()
}
