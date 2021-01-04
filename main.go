package main

import (
	"Information_Retrieval/storage"
	"log"
)

func main() {
	//i := index.NewIndex("./p_docs", 6)
	//indexFile := i.Construct()
	//v := vector_space.NewVectorizer(indexFile, 3)
	//v.Vectorize()
	//v.Query("نشست کمیسیون")

	//c := championlist.NewChampion("./blocks2/1.txt", 1)
	//c.Create()

	log.Printf("%#v\n", storage.NewMinioConnection()) // minioClient is now setup
	//fmt.Println(query.Query())
}
