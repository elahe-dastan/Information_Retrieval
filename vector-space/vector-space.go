package vector_space

import (
	"Information_Retrieval/tokenize"
	"io/ioutil"
	"log"
	"math"
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
	v.idf = make([]float64, len(lines)-1)

	for i, l := range lines{
		if l == "" {
			continue
		}

		termPostingList := tokenize.Unmarshal(l)
		count := 1
		for j := 1; j < len(termPostingList.PostingList); j++ {
			if termPostingList.PostingList[j] != termPostingList.PostingList[j - 1]{
				count++
			}
		}

		v.idf[i] = math.Log10(float64(v.docsNum / count))
	}
}

func (v *Vectorizer) calculateTF() {
	
}
