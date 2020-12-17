package bsbi

type Bsbi struct {
	blockNum int
}

func NewBsbi() *Bsbi {
	return &Bsbi{blockNum: 0}
}

func (b *Bsbi) Block(){
	print("raha")
	//b.blockNum++
	//
	//m := make(map[string]int)
	//for {
	//	// read 160 bytes
	//
	//	sortedBlock := BlockSort(termDocs)
	//	sortedStringInMemory := ""
	//	dir := "/home/raha/go/src/Information_Retrieval/blocks0"
	//
	//	err = os.Mkdir(dir, 0700)
	//	//if err != nil {
	//	//	if err != os.
	//	//	log.Fatal(err)
	//	//}
	//
	//	dir += "/" + strconv.Itoa(block) + ".txt"
	//	err = os.Chmod(dir, 0700)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	o, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	previous := TermDoc{
	//		term: "",
	//		doc:  0,
	//	}
	//	for i := range sortedBlock {
	//		sb := sortedBlock[i]
	//		if sb == previous {
	//			continue
	//		}
	//
	//		if sb.term == previous.term {
	//			sortedStringInMemory += "," + strconv.Itoa(sb.doc)
	//			continue
	//		}
	//
	//		if previous.term != "" {
	//			sortedStringInMemory += "\n"
	//		}
	//		sortedStringInMemory += strconv.Itoa(sb.doc) + " " + sb.term
	//		previous = sb
	//	}
	//
	//	_, err = o.WriteString(sortedStringInMemory)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	//err = os.Chmod("./tokens.txt", 0777)
	//	//if err != nil {
	//	//	fmt.Println(err)
	//	//}
	//	//
	//	//o, err = os.OpenFile("./tokens.txt",os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	//	//
	//	//// using buffer ################################
	//	//_, err = o.WriteString(stringInMemory)
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	//	block++
	//}
}