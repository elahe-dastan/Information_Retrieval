package tokenize

import (
	"bufio"
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

type TermPostingList struct {
	term        string
	postingList []int // always sorted
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
		dir := "/home/raha/go/src/Information_Retrieval/blocks0"

		err = os.Mkdir(dir, 0700)
		//if err != nil {
		//	if err != os.
		//	log.Fatal(err)
		//}

		dir += "/" + strconv.Itoa(block) + ".txt"
		err = os.Chmod(dir, 0700)
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
	merge()
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

func merge() {
	mergeRun := 0
	for {
		// all blocks
		blocks, err := ioutil.ReadDir("./blocks" + strconv.Itoa(mergeRun))
		if err != nil {
			log.Fatal(err)
		}

		if len(blocks) == 1 {
			break
		}

		blockNames := make([]string, len(blocks))

		for i, b := range blocks {
			// open each doc and tokenize
			blockNames[i] = b.Name()
		}

		// open up to 10 files
		start := 0
		size := 10
		if len(blockNames) < 10 {
			size = len(blockNames)
		}
		end := start + size
		filePointers := make([]*bufio.Scanner, size)
		termPostingLists := make(PairArrayList, size)
		outputBuffer := make([]byte, 160)
		newBlocks := 0

		dir := "/home/raha/go/src/Information_Retrieval/blocks" + strconv.Itoa(mergeRun+1)

		err = os.Mkdir(dir, 0700)
		//if err != nil {
		//	if err != os.
		//	log.Fatal(err)
		//}

		// MERGE RUN
		for {
			dir += "/" + strconv.Itoa(newBlocks) + ".txt"

			newBlocks++
			err = os.Chmod(dir, 0700)
			if err != nil {
				fmt.Println(err)
			}

			//output file
			o, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE, os.ModeAppend)
			if err != nil {
				log.Fatal(err)
			}

			residual := 160
			for i := start; i < end; i++ {
				f, err := os.Open("./blocks" + strconv.Itoa(mergeRun) + "/" + blockNames[i])
				if err != nil {
					log.Fatal(err)
				}

				scanner := bufio.NewScanner(f)
				scanner.Split(bufio.ScanLines)
				filePointers[i] = scanner
			}
			// 10 files
			for {
				// how to move pointer forward
				for i := 0; i < size; i++ {
					s := filePointers[i]

					if !s.Scan() {
						break
					}

					l := strings.Split(s.Text(), "")
					postingList := make([]int, 0)
					postings := strings.Split(l[1], ",")
					for j := 0; j < len(postings); j++ {
						posting, err := strconv.Atoi(postings[j])
						if err != nil {
							log.Fatal()
						}

						postingList = append(postingList, posting)
					}

					termPostingLists[i] = PairArray{
						Key:   l[0],
						Value: postingList,
					}

				}

				sort.Sort(sort.Reverse(termPostingLists))

				firstTerm := termPostingLists[0].Key
				firstPostingList := termPostingLists[0].Value

				for i := 1; i < size; i++ {
					if termPostingLists[i].Key != firstTerm {
						break
					}

					for j := 0; i < len(termPostingLists[i].Value); j++ {
						firstPostingList = append(firstPostingList, termPostingLists[i].Value[j])
					}
				}

				sort.Ints(firstPostingList)
				firstPostingListStr := make([]string, len(firstPostingList))
				for k := 0; k < len(firstPostingList); k++ {
					firstPostingListStr = append(firstPostingListStr, strconv.Itoa(firstPostingList[k]))
				}

				out := firstTerm + " " + strings.Join(firstPostingListStr, ",")

				outBytes := []byte(out)

				if len(outBytes) < residual {
					outputBuffer = append(outputBuffer, outBytes...)
					residual -= len(outBytes)
				} else {
					outputBuffer = append(outputBuffer, outBytes[:residual]...)
					_, err := o.Write(outputBuffer)
					if err != nil {
						log.Fatal(err)
					}
					outBytes = outBytes[residual:]

					_, err = o.Write(outputBuffer)
					if err != nil {
						log.Fatal(err)
					}

				}
			}
		}

		mergeRun++
	}
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Key < p[j].Key }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type PairArray struct {
	Key   string
	Value []int
}

type PairArrayList []PairArray

func (p PairArrayList) Len() int           { return len(p) }
func (p PairArrayList) Less(i, j int) bool { return p[i].Key < p[j].Key }
func (p PairArrayList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
