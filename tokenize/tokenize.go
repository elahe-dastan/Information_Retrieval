package tokenize

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type TermDoc struct {
	term string // 16 byte
	doc  int // 8 byte
}

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
			termDocs = append(termDocs, TermDoc{
				term: token,
				doc:  id,
			})
		}


	}
}
