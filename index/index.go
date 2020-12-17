package index

import (
	"Information_Retrieval/bsbi"
	"Information_Retrieval/tokenize"
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

type index struct {
	collectionDir string
	memorySize    int
	docId         int
	sortAlgorithm *bsbi.Bsbi
}

func NewIndex(collectionDir string, memorySize int) *index {
	return &index{collectionDir: collectionDir, memorySize: memorySize, docId: 0, sortAlgorithm: bsbi.NewBsbi("./blocks/")}
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

	docDir := i.collectionDir + "/" + docName

	f, err := os.Open(docDir)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	i.tokenizeSortBlock(f)
}

func (i *index) tokenizeSortBlock(f *os.File) {
	memIndex := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	termDocs := make([]tokenize.TermDoc, i.memorySize)
	for scanner.Scan() {
		term := scanner.Text()
		termDocs[memIndex] = tokenize.TermDoc{
			Term: term,
			Doc:  i.docId,
		}

		memIndex++
		if memIndex == i.memorySize {
			i.sortAlgorithm.WriteBlock(termDocs)
			termDocs = make([]tokenize.TermDoc, i.memorySize)
			memIndex = 0
		}
	}

	if len(termDocs) > 0 {
		i.sortAlgorithm.WriteBlock(termDocs)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
