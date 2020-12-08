package sort

import (
	"log"
	"os"
)



func Bsbi()  {
	o, err := os.Open("./tokens.txt")
	if err != nil {
		log.Fatal(err)
	}

	for {
		mem := make([]byte, 160)
		_, err = o.Read(mem)
		if err != nil {
			log.Fatal(err)
		}
	}

}