package main

import (
	"Information_Retrieval/index"
	"Information_Retrieval/storage"
)

func main() {
	minioClient := storage.NewMinioConnection() // minioClient is now setup
	_ = index.NewIndex(minioClient, 6)
	//indexFile := i.Construct()
	//v := vector_space.NewVectorizer(indexFile, 3)
	//v.Vectorize()
	//v.Query("نشست کمیسیون")

	//c := championlist.NewChampion("./blocks2/1.txt", 1)
	//c.Create()

	//fmt.Println(query.Query())
}
