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



type TermPostingList struct {
	term        string
	postingList []int // always sorted
}

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
			// open each doc and tokeni
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

			for i := 0; i < size; i++ {
				s := filePointers[i]

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

				termPostingLists[i] = Head{
					Pointer: s,
					Term: TermPostingList{
						term:        l[0],
						postingList: postingList,
					},
				}

			}

			sort.Sort(sort.Reverse(termPostingLists))

			// 10 files
			for {
				// how to move pointer forward
				firstTerm := termPostingLists[0].Term.term
				firstPostingList := termPostingLists[0].Term.postingList

				firstPointer := termPostingLists[0].Pointer

				if !firstPointer.Scan() {
					termPostingLists = termPostingLists[1:]
				}else {
					l := strings.Split(firstPointer.Text(), "")
					postingList := make([]int, 0)
					postings := strings.Split(l[1], ",")
					for j := 0; j < len(postings); j++ {
						posting, err := strconv.Atoi(postings[j])
						if err != nil {
							log.Fatal()
						}

						postingList = append(postingList, posting)
					}

					termPostingLists[0] = Head {
						Pointer: firstPointer,
						Term:    TermPostingList{
							term:        l[0],
							postingList: postingList,
						},
					}
				}

				for i := 1; i < size; i++ {
					if termPostingLists[i].Term.term != firstTerm {
						break
					}

					firstPostingList = append(firstPostingList, termPostingLists[i].Term.postingList...)
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

type Head struct {
	Pointer *bufio.Scanner
	Term    TermPostingList
}

type PairArrayList []Head

func (p PairArrayList) Len() int           { return len(p) }
func (p PairArrayList) Less(i, j int) bool { return p[i].Term.term < p[j].Term.term }
func (p PairArrayList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
