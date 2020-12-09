package tokenize

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TermDoc struct {
	term string // 16 byte
	doc  int    // 8 byte
}

var m map[string]int

const stopWordCount = 4

// our memory can keep only 240 bytes
//var termDocs [10]TermDoc

func AllDocs() {
	// go through docs in a specified directory
	docs, err := ioutil.ReadDir("./docs")
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range docs {
		// open each doc and tokenize
		doc(d.Name())
	}
}

// Yes, I know we have few small docs but our logic should work for huge ones too so to stimulate the situation we think
// our memory is too small
func doc(name string) {
	id := strings.Split(name, ".")[0]

	f, err := os.Open("./docs/" + name)
	if err != nil {
		log.Fatal(err)
	}

	block := 0
	m = make(map[string]int)
	for {
		// read 160 bytes
		mem := make([]byte, 160)
		n, err := f.Read(mem)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}

		s := string(mem[:n])
		tokens := strings.Fields(s)
		var termDocs []TermDoc
		for _, token := range tokens {
			m[token]++
			idInt, _ := strconv.Atoi(id)
			termDocs = append(termDocs, TermDoc{
				term: token,
				doc:  idInt,
			})
		}

		sortedBlock := BlockSort(termDocs)
		sortedStringInMemory := ""
		dir := "./blocks/" + strconv.Itoa(block) + ".txt"
		err = os.Chmod(dir, 0777)
		if err != nil {
			fmt.Println(err)
		}

		o, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE, os.ModeAppend)
		if err != nil {
			log.Fatal(err)
		}

		previous := TermDoc{
			term: "",
			doc:  0,
		}
		for i := range sortedBlock {
			sb := sortedBlock[i]
			if sb == previous {
				continue
			}

			if sb.term == previous.term {
				sortedStringInMemory += "," + strconv.Itoa(sb.doc)
				continue
			}

			if previous.term != "" {
				sortedStringInMemory += "\n"
			}
			sortedStringInMemory += strconv.Itoa(sb.doc) + " " + sb.term
			previous = sb
		}

		_, err = o.WriteString(sortedStringInMemory)
		if err != nil {
			log.Fatal(err)
		}

		//err = os.Chmod("./tokens.txt", 0777)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//
		//o, err = os.OpenFile("./tokens.txt",os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//// using buffer ################################
		//_, err = o.WriteString(stringInMemory)
		//if err != nil {
		//	log.Fatal(err)
		//}
		block++
	}

	//stopWords := stopWord()
	//merge
}

func BlockSort(termDocs []TermDoc) []TermDoc {
	if len(termDocs) < 2 {
		return termDocs
	}

	left, right := 0, len(termDocs)-1

	pivot := rand.Int() % len(termDocs)

	termDocs[pivot], termDocs[right] = termDocs[right], termDocs[pivot]

	for i, _ := range termDocs {
		if termDocs[i].term < termDocs[right].term {
			termDocs[left], termDocs[i] = termDocs[i], termDocs[left]
			left++
		}
	}

	termDocs[left], termDocs[right] = termDocs[right], termDocs[left]

	BlockSort(termDocs[:left])
	BlockSort(termDocs[left+1:])

	return termDocs

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
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
