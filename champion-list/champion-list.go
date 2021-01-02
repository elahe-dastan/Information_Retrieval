package champion_list

import (
	"Information_Retrieval/tokenize"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type champion struct {
	termPostingLists []tokenize.TermPostingList
	championFile *os.File
}

func NewChampion(indexFile string) *champion {
	dat, err := ioutil.ReadFile(indexFile)
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

	championFile := "./champion.txt"
	o, err := os.OpenFile(championFile, os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chmod(championFile, 0700)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(termPostingLists)

	return &champion{
		termPostingLists: termPostingLists,
		championFile: o,
	}
}

func Create(indexFile string) {
	//h := &Heap{}
	//heap.Init(h)
}
