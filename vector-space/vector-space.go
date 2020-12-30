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
	termIndex        map[string]int
}

func NewVectorizer(indexPath string, docsNum int) *Vectorizer {
	dat, err := ioutil.ReadFile(indexPath)
	if err != nil {
		log.Fatal(err)
	}

	tmp := strings.Split(string(dat), "\n")
	lines := tmp[:len(tmp)-1]
	termPostingLists := make([]tokenize.TermPostingList, len(lines))

	termIndex := make(map[string]int)

	for i, l := range lines {
		termPostingList := tokenize.Unmarshal(l)
		termPostingLists[i] = termPostingList
		termIndex[termPostingList.Term] = i
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
		termIndex:        termIndex,
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
			v.tfIdf[i][j] = (math.Log10(1 + float64(v.tf[i][j]))) * v.idf[j]
		}
	}
}

func (v *Vectorizer) Query(query string) {
	queryVector := v.queryVectorizer(query)
}

func (v *Vectorizer) queryVectorizer(query string) []float64 {
	vector := make([]float64, v.termsNum)

	tokens := strings.Split(query, " ")
	for _, t :=  range tokens{
		index, ok := v.termIndex[t]
		if !ok{
			continue
		}
		vector[index]++
	}

	return vector
}

func (v *Vectorizer) cosineSimilarity(queryVector []float64) {
	
}
