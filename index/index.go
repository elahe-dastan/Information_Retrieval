package index

import (
	"Information_Retrieval/bsbi"
	"Information_Retrieval/normalize"
	"Information_Retrieval/tokenize"
	"bufio"
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"os"
	"strconv"
)

type index struct {
	minioClient   *minio.Client
	memorySize    int
	docId         int
	sortAlgorithm *bsbi.Bsbi
}

func NewIndex(minioClient *minio.Client, memorySize int) *index {
	return &index{minioClient: minioClient, memorySize: memorySize, docId: 0, sortAlgorithm: bsbi.NewBsbi(10, memorySize)}
}

// dir is document collection directory
func (i *index) Construct() {
	found, err := i.minioClient.BucketExists(context.Background(), "information-retrieval")
	if err != nil {
		log.Fatal(err)
	}
	if !found {
		log.Fatal("information-retrieval bucket not found")
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := i.minioClient.ListObjects(ctx, "information-retrieval", minio.ListObjectsOptions{})
	for object := range objectCh {
		if object.Err != nil {
			log.Fatal(object.Err)
		}
		fmt.Println(object.Size)
	}


	//docs, err := ioutil.ReadDir(i.minioClient)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, d := range docs {
	//	i.construct(d.Name())
	//}
	//
	//return i.sortAlgorithm.Merge()
}

// construct index for one document
func (i *index) construct(docName string) {
	//i.docId++
	//
	//docDir := i.minioClient + "/" + docName
	//
	//f, err := os.Open(docDir)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer f.Close()
	//
	//i.tokenizeSortBlock(f)
}

func (i *index) tokenizeSortBlock(f *os.File) {
	memIndex := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	termPostingList := make([]tokenize.TermPostingList, 0)
	for scanner.Scan() {
		word := scanner.Text()
		terms := normalize.Normalize(word)
		for _, term := range terms {
			termPostingList = append(termPostingList, tokenize.TermPostingList{
				Term:        term,
				PostingList: []string{strconv.Itoa(i.docId)},
			})

			memIndex++
		}

		if memIndex >= i.memorySize {
			i.sortAlgorithm.WriteBlock(termPostingList)
			termPostingList = make([]tokenize.TermPostingList, 0)
			memIndex = 0
		}
	}

	// masmali
	a := make([]tokenize.TermPostingList, 0)
	for i := range termPostingList {
		if termPostingList[i].Term == "" {
			break
		}

		a = append(a, termPostingList[i])
	}

	if len(termPostingList) > 0 {
		i.sortAlgorithm.WriteBlock(a)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
