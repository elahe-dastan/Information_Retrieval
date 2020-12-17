package bsbi

import (
	"Information_Retrieval/tokenize"
	"bufio"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Bsbi struct {
	blockDir       string
	openFileNum    int
	outPutBuffSize int
	blockNum       int
	mergeRun       int
}

func NewBsbi(blockDir string, openFilesNum int, outPutBuffSize int) *Bsbi {
	err := os.Mkdir(blockDir+"0", 0700)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	return &Bsbi{blockDir: blockDir, openFileNum: openFilesNum, outPutBuffSize: outPutBuffSize, blockNum: 0, mergeRun: 0}
}

func (b *Bsbi) WriteBlock(termDocs []tokenize.TermDoc) {
	b.blockNum++

	sortedBlock := sortBlock(termDocs)

	filePath := b.blockDir + strconv.Itoa(b.blockNum) + ".txt"
	o, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chmod(filePath, 0700)
	if err != nil {
		log.Fatal(err)
	}

	sortedBlockStr := ""

	var previous tokenize.TermDoc
	for i := range sortedBlock {
		termDoc := sortedBlock[i]
		if termDoc == previous {
			continue
		}

		if termDoc.Term == previous.Term {
			sortedBlockStr += "," + strconv.Itoa(termDoc.Doc)
			continue
		}

		if previous.Term != "" {
			sortedBlockStr += "\n"
		}
		sortedBlockStr += termDoc.Term + " " + strconv.Itoa(termDoc.Doc)
		previous = termDoc
	}

	_, err = o.WriteString(sortedBlockStr)
	if err != nil {
		log.Fatal(err)
	}

}

func sortBlock(termDocs []tokenize.TermDoc) []tokenize.TermDoc {
	if len(termDocs) < 2 {
		return termDocs
	}

	left, right := 0, len(termDocs)-1

	pivot := rand.Int() % len(termDocs)

	termDocs[pivot], termDocs[right] = termDocs[right], termDocs[pivot]

	for i, _ := range termDocs {
		if termDocs[i].Term < termDocs[right].Term {
			termDocs[left], termDocs[i] = termDocs[i], termDocs[left]
			left++
		}
	}

	termDocs[left], termDocs[right] = termDocs[right], termDocs[left]

	sortBlock(termDocs[:left])
	sortBlock(termDocs[left+1:])

	return termDocs
}

func (b *Bsbi) Merge() {
	mergeRun := 0
	for {
		// all blocks
		blocks, err := ioutil.ReadDir(b.blockDir + strconv.Itoa(mergeRun))
		if err != nil {
			log.Fatal(err)
		}

		mergeRun++

		if len(blocks) < b.openFileNum {
			lastMerge()
			break
		}

		b.middleMerge(blocks)
	}
}

func (b *Bsbi) middleMerge(blocks []os.FileInfo) {
	block := 0
	blockNames := make([]string, len(blocks))

	for i, b := range blocks {
		blockNames[i] = b.Name()
	}
	//if len(blockNames) < 10 {
	//	size = len(blockNames)
	//}
	//end := start + size
	filePointers := make([]*bufio.Scanner, b.openFileNum)
	for i := 0; i < b.openFileNum; i++ {
		f, err := os.Open(b.blockDir + blockNames[i]) // it may need a / in between
		defer f.Close()

		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		filePointers[i] = scanner
	}

	outputBuffer := make([]tokenize.TermPostingList, b.outPutBuffSize)
	outputDir := b.blockDir + strconv.Itoa(b.mergeRun)
	err := os.Mkdir(outputDir, 0700)
	if err != nil && !os.IsExist(err){
		log.Fatal(err)
	}

	//output file
	o, err := os.OpenFile(outputDir + strconv.Itoa(block) + ".txt", os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	termPostingLists := make(PairArrayList, b.openFileNum)

	residual := 160

	for i := 0; i < b.openFileNum; i++ {
		s := filePointers[i]

		// write marshal unmarshal for it
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
		} else {
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

			termPostingLists[0] = Head{
				Pointer: firstPointer,
				Term: TermPostingList{
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

}
