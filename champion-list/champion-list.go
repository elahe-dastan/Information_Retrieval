package champion_list

import (
	"Information_Retrieval/tokenize"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type frequency struct {
	docId string
	freq  int
}

type champion struct {
	termPostingLists []tokenize.TermPostingList
	championFile     *os.File
	k                int
}

func NewChampion(indexFile string, k int) *champion {
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

	return &champion{
		termPostingLists: termPostingLists,
		championFile:     o,
		k:                k,
	}
}

func (c *champion) Create() {
	for _, t := range c.termPostingLists {
		frequencies := make([]frequency, 0)
		previous := "0"
		for _, docId := range t.PostingList {
			if docId != previous {
				frequencies = append(frequencies, frequency{
					docId: docId,
					freq:  1,
				})
				previous = docId
			} else {
				p := frequencies[len(frequencies)-1]
				p.freq++
				frequencies[len(frequencies)-1] = p
			}
		}
		fmt.Println(frequencies)
	}
	//h := &Heap{}
	//heap.Init(h)
	//_, err = o.WriteString(sortedBlockStr)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

type Frequencies []frequency

func (f Frequencies) Len() int           { return len(f) }
func (f Frequencies) Less(i, j int) bool { return f[i].freq < f[j].freq }
func (f Frequencies) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
