package vector_space

import (
	"Information_Retrieval/tokenize"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

type Vectorizer struct {
	termPostingLists []tokenize.TermPostingList
	docsNum          int
	termsNum         int
	idf              []float64
}

func NewVectorizer(indexPath string, docsNum int) *Vectorizer {
	dat, err := ioutil.ReadFile(indexPath)
	if err != nil {
		log.Fatal(err)
	}

	tmp := strings.Split(string(dat), "\n")
	lines := tmp[:len(tmp)-1]
	termPostingLists := make([]tokenize.TermPostingList, len(lines))

	for i, l := range lines {
		termPostingList := tokenize.Unmarshal(l)
		termPostingLists[i] = termPostingList
	}

	return &Vectorizer{
		termPostingLists: termPostingLists,
		docsNum:          docsNum,
		termsNum:         len(lines),
	}
}

func (v *Vectorizer) Vectorize() {
	v.calculateIDF()

}

func (v *Vectorizer) calculateIDF() {
	v.idf = make([]float64, v.termsNum)

	for i, t := range v.termPostingLists {
		count := 1
		for j := 1; j < len(t.PostingList); j++ {
			if t.PostingList[j] != t.PostingList[j-1] {
				count++
			}
		}

		v.idf[i] = math.Log10(float64(v.docsNum / count))
	}

	fmt.Println(v.idf)
}

func (v *Vectorizer) calculateTF() {

}
