package vector_space

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Vectorizer struct {
	indexPath string
	docsNum   int
	IDF       []float64
}

func NewVectorizer(indexPath string, docsNum int) *Vectorizer {
	return &Vectorizer{
		indexPath: indexPath,
		docsNum:   docsNum,
		IDF:       make([]float64, docsNum),
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

	fmt.Print(string(dat))
}
