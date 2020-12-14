package index

import (
	"io/ioutil"
	"log"
)

// dir is document collection directory
func Construct(dir string) {
	docs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range docs {
		// open each doc and tokenize
		doc(d.Name())
	}
}
