package tokenize

type TermDoc struct {
	Term string
	Doc  int
}

type TermPostingList struct {
	term        string
	postingList []int // always sorted
}
