package oracle

import (
	"math/rand/v2"

	"401/automata"
)

// Gen is a seeded generator of random words and random ASTs.
//
// It is seeded rather than global on purpose: a failing property test is only
// useful if you can rerun it, and a shared global source makes the failing
// input depend on test ordering.
type Gen struct {
	rng      *rand.Rand
	alphabet []automata.Symbol
}

// NewGen returns a generator over alphabet, deterministic in seed.
func NewGen(seed uint64, alphabet []automata.Symbol) *Gen {
	if len(alphabet) == 0 {
		panic("oracle: NewGen needs a non-empty alphabet")
	}
	return &Gen{
		rng:      rand.New(rand.NewPCG(seed, seed^0x9e3779b97f4a7c15)),
		alphabet: alphabet,
	}
}

// Word returns a random word of length in [0, maxLen].
//
// Length zero is included deliberately: the empty word is where off-by-one
// errors in Run, in epsilon-closures, and in Thompson fragments show up.
func (g *Gen) Word(maxLen int) automata.Word {
	n := g.rng.IntN(maxLen + 1)
	w := make(automata.Word, n)
	for i := range w {
		w[i] = g.alphabet[g.rng.IntN(len(g.alphabet))]
	}
	return w
}

// Words returns n random words, each of length in [0, maxLen].
func (g *Gen) Words(n, maxLen int) []automata.Word {
	ws := make([]automata.Word, n)
	for i := range ws {
		ws[i] = g.Word(maxLen)
	}
	return ws
}
