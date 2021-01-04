package main

import champion_list "Information_Retrieval/champion-list"

func main() {
	//i := index.NewIndex("./p_docs", 6)
	//indexFile := i.Construct()
	//v := vector_space.NewVectorizer(indexFile, 3)
	//v.Vectorize()
	//v.Query("نشست کمیسیون")

	c := champion_list.NewChampion("./blocks2/1.txt", 1)
	c.Create()

	//fmt.Println(query.Query())
}
