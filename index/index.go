package index

import (
	"Information_Retrieval/bsbi"
	"io/ioutil"
	"log"
	"os"
)

type index struct {
	collectionDir string
	docId         int
	sortAlgorithm *bsbi.Bsbi
}

func NewIndex(collectionDir string) *index {
	return &index{collectionDir: collectionDir, docId: 0, sortAlgorithm: bsbi.NewBsbi()}
}

// dir is document collection directory
func (i *index) Construct() {
	docs, err := ioutil.ReadDir(i.collectionDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range docs {
		i.construct(d.Name())
	}
}

// construct index for one document
func (i *index) construct(docName string) {
	i.docId++

	f, err := os.Open(i.collectionDir + "/" + docName)
	if err != nil {
		log.Fatal(err)
	}

	i.sortAlgorithm.Block()
}
