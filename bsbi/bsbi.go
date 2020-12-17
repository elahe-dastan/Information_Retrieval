package bsbi

import (
	"Information_Retrieval/tokenize"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type Bsbi struct {
	blockDir string
	blockNum int
}

func NewBsbi(blockDir string) *Bsbi {
	err := os.Mkdir(blockDir, 0700)
	if err != nil {
		log.Fatal(err)
	}

	return &Bsbi{blockDir: blockDir, blockNum: 0}
}

func (b *Bsbi) WriteBlock(termDocs []tokenize.TermDoc){
	b.blockNum++

	for {
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