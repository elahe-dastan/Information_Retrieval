package vector_space

import (
	"Information_Retrieval/tokenize"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Vectorizer struct {
	termPostingLists []tokenize.TermPostingList
	docsNum          int
	termsNum         int
	idf              []float64
	tf               [][]int
	tfIdf            [][]float64
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

	tf := make([][]int, docsNum)
	for i := 0; i < docsNum; i++ {
		tf[i] = make([]int, len(lines))
	}

	tfIdf := make([][]float64, docsNum)
	for i := 0; i < docsNum; i++ {
		tfIdf[i] = make([]float64, len(lines))
	}

	return &Vectorizer{
		termPostingLists: termPostingLists,
		docsNum:          docsNum,
		termsNum:         len(lines),
		tf:               tf,
		tfIdf:            tfIdf,
	}
}

func (v *Vectorizer) Vectorize() {
	v.calculateIDF()
	v.calculateTF()
	v.calculateTFIDF()
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
}

func (v *Vectorizer) calculateTF() {
	// i expresses term index
	for i, t := range v.termPostingLists {
		for j := 0; j < len(t.PostingList); j++ {
			docId, err := strconv.Atoi(t.PostingList[j])
			if err != nil {
				log.Fatal(err)
			}

			v.tf[docId-1][i]++
		}
	}
}

func (v *Vectorizer) calculateTFIDF() {
	for i := 0; i < v.docsNum; i++ {
		for j := 0; j < v.termsNum; j++ {
			v.tfIdf[i][j] = (1 + math.Log10(float64(v.tf[i][j]))) * v.idf[j]
		}
	}
}
