package vector_space

import (
	"Information_Retrieval/tokenize"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Vectorizer struct {
	indexPath string
	docsNum   int
	idf       []float64
}

func NewVectorizer(indexPath string, docsNum int) *Vectorizer {
	return &Vectorizer{
		indexPath: indexPath,
		docsNum:   docsNum,
		idf:       nil,
	}
}

func (v *Vectorizer) Vectorize() {
	v.calculateIDF()
}

func (v *Vectorizer) calculateIDF() {
	dat, err := ioutil.ReadFile(v.indexPath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(dat), "\n")
	v.idf = make([]float64, len(lines))

	for i, l := range lines{
		termPostingList := tokenize.Unmarshal(l)
		v.idf[i] = float64(len(termPostingList.PostingList))
		fmt.Println(v.idf)
	}
}
