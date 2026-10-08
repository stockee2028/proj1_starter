package automata

// FiniteLanguage is a language given as an explicit, finite list of its words.
//
// The contract fixes the name, NewFiniteLanguage, and Contains, not what is
// inside: a map, a sorted slice, a trie, or a plain []Word are all fine. Its
// words may contain any bytes; the ASCII letter/digit restriction applies only
// to machine alphabets.
type FiniteLanguage struct {
	// your representation here
}

// NewFiniteLanguage returns the language whose members are exactly the given
// words. Repeated words are fine and change nothing. With no arguments it returns
// the empty language. There is no error return because a list of words cannot
// be malformed.
func NewFiniteLanguage(words ...Word) *FiniteLanguage {
	panic("unimplemented: phase 1.0 NewFiniteLanguage")
}

// Contains reports whether w is exactly one of the language's words. A nil w
// is the empty word.
func (l *FiniteLanguage) Contains(w Word) bool {
	panic("unimplemented: phase 1.0 FiniteLanguage.Contains")
}
