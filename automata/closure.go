package automata

// Phase 1.2 — closure operations, and the decision procedures built from them: the
// product construction (Sipser Theorem 1.25) and E_DFA and EQ_DFA (Sipser
// Theorems 4.4 and 4.5).

// Intersect returns a DFA for L(a) intersect L(b), by the product construction.
//
// a and b must have the same alphabet, and the result has that alphabet too.
// Alphabets are compared as sets: neither the order of the symbols passed to
// NewDFA nor repeats among them matter, so machines built over {0,1} and
// {1,0,1} have the same alphabet. If the alphabets differ as sets, even when
// one contains the other, Intersect returns an error.
func Intersect(a, b *DFA) (*DFA, error) {
	panic("unimplemented: phase 1.2 Intersect")
}

// Union returns a DFA for L(a) union L(b).
//
// Alphabets work as in Intersect: a and b must have the same alphabet,
// compared as sets, and Union returns an error if they don't.
func Union(a, b *DFA) (*DFA, error) {
	panic("unimplemented: phase 1.2 Union")
}

// Complement returns a DFA for the complement of L(a), over a's alphabet.
//
// The complement is relative to Sigma*, where Sigma is a's alphabet: a word
// containing a symbol outside Sigma is in neither L(a) nor its complement.
func Complement(a *DFA) *DFA {
	panic("unimplemented: phase 1.2 Complement")
}

// Empty reports whether L(a) is the empty language.
func Empty(a *DFA) bool {
	panic("unimplemented: phase 1.2 Empty")
}

// Equivalent reports whether L(a) equals L(b), by deciding whether their
// symmetric difference is empty:
//
//	Equivalent(a, b)  <=>  Empty(Intersect(a, Complement(b)))
//	                  and  Empty(Intersect(Complement(a), b))
//
// It panics if a and b have different alphabets, since it has no error return
// to report one with. It returns false exactly when Witness(a, b) reports a
// word.
func Equivalent(a, b *DFA) bool {
	panic("unimplemented: phase 1.2 Equivalent")
}

// Witness returns a shortest word on which a and b disagree, and true, or the
// zero Word and false if L(a) equals L(b).
//
// Shortest means by length; when several words tie, any of them will do. The
// empty word can be a witness, so check the bool, not the word.
//
// Witness panics if a and b have different alphabets. It reports a word exactly
// when Equivalent(a, b) is false.
func Witness(a, b *DFA) (Word, bool) {
	panic("unimplemented: phase 1.2 Witness")
}
