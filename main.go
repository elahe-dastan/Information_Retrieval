package main

import "Information_Retrieval/index"

func main() {
	i := index.NewIndex("./docs")
	i.Construct()
}
