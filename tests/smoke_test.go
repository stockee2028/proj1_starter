package tests

import (
	"testing"

	"401/automata"
	"401/corpus"
	"401/oracle"
)

// TestP1_0_Types passes on a fresh clone, before you have written anything: it
// exercises only automata/lang.go, which ships complete. make smoke runs it.
// The tests after it fail with "unimplemented: phase 1.x ..." until you build the
// phase in question.

// evenLength is a hand-written Language: anything with a Contains method is
// one.
type evenLength struct{}

func (evenLength) Contains(w automata.Word) bool { return len(w)%2 == 0 }

func TestP1_0_Types(t *testing.T) {
	// Word conversion from a Go string gives you bytes, which is the whole
	// reason Symbol is byte.
	if got := automata.Word("0011"); len(got) != 4 || got[0] != '0' || got[3] != '1' {
		t.Fatalf("automata.Word(%q) = %v; want the four bytes of the string", "0011", got)
	}

	var l automata.Language = evenLength{}
	oracle.Accepts(t, l, "")
	oracle.Accepts(t, l, "0011")
	oracle.Rejects(t, l, "001")
}

// TestP1_0_FiniteLanguage is the first test that needs your code:
// NewFiniteLanguage and Contains.
func TestP1_0_FiniteLanguage(t *testing.T) {
	oracle.Case(t, "Empty", func(t *testing.T) {
		empty := automata.NewFiniteLanguage()
		oracle.Rejects(t, empty, "")
		oracle.Rejects(t, empty, "ab")
	})

	oracle.Case(t, "Membership", func(t *testing.T) {
		l := automata.NewFiniteLanguage(automata.Word(""), automata.Word("ab"), automata.Word("ba"))
		oracle.Accepts(t, l, "") // the empty word is a member here
		oracle.Accepts(t, l, "ab")
		oracle.Accepts(t, l, "ba")
		oracle.Rejects(t, l, "a")
		oracle.Rejects(t, l, "abab")
	})
}

// TestP1_1_Corpus is the first test that needs an automaton: NewDFA, Run, and
// Contains.
func TestP1_1_Corpus(t *testing.T) {
	oracle.Case(t, "EvenZeros", func(t *testing.T) {
		d := corpus.EvenZeros()
		oracle.Accepts(t, d, "")     // zero 0s is even
		oracle.Accepts(t, d, "1001") // two
		oracle.Rejects(t, d, "0")    // one
		oracle.Rejects(t, d, "0111") // one
	})

	oracle.Case(t, "EndsWith01", func(t *testing.T) {
		d := corpus.EndsWith01()
		oracle.Accepts(t, d, "01")
		oracle.Accepts(t, d, "110101")
		oracle.Rejects(t, d, "")
		oracle.Rejects(t, d, "010")
	})

	oracle.Case(t, "DivisibleBy3", func(t *testing.T) {
		d := corpus.DivisibleBy3()
		oracle.Accepts(t, d, "11")   // 3
		oracle.Accepts(t, d, "110")  // 6
		oracle.Accepts(t, d, "1001") // 9
		oracle.Rejects(t, d, "1")    // 1
		oracle.Rejects(t, d, "100")  // 4
	})
}

// TestP1_2_Closure is the worked example from the README: a language under
// test written as a closure expression over the corpus.
func TestP1_2_Closure(t *testing.T) {
	oracle.Case(t, "even zeros and not ending in 01", func(t *testing.T) {
		lang := oracle.Must(automata.Intersect(
			corpus.EvenZeros(), automata.Complement(corpus.EndsWith01())))

		oracle.Accepts(t, lang, "0011") // two 0s, and does not end in 01
		oracle.Rejects(t, lang, "0101") // two 0s, but ends in 01 — so the
		//                                 Complement half is what rejects it
	})
}
