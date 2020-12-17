package tokeni

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)





var m map[string]int

const stopWordCount = 4

//var termDocs [10]TermDoc


//Yes, I know we have few small docs but our logic should work for huge ones too so to stimulate the situation we think
//our memory is too small
func doc(name string) {


	//stopWords := stopWord()
	merge()
}



func stopWord() []string {
	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	stopWords := make([]string, stopWordCount)
	for i := 0; i < stopWordCount; i++ {
		stopWords[i] = pl[i].Key
	}

	return stopWords
}


type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Key < p[j].Key }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Head struct {
	Pointer *bufio.Scanner
	Term    TermPostingList
}

type PairArrayList []Head

func (p PairArrayList) Len() int           { return len(p) }
func (p PairArrayList) Less(i, j int) bool { return p[i].Term.term < p[j].Term.term }
func (p PairArrayList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
