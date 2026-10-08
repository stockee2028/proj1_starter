package engine

import (
	"401/automata"
	"401/regex"
)

// Phase 1.9 — the engine.

// Regexp is a compiled regular expression, ready to match any number of words.
//
// What you store here is the design decision described on FromAST.
type Regexp struct {
	// your representation here
}

// FromAST returns a Regexp matching the language of n. The basic wiring is
// Node -> regex.Compile -> automata.NFA -> Run. It leaves n unchanged.
//
// There is a design decision here, and the handout asks you to justify and
// benchmark it. Given an AST, you can either:
//
//   - keep the Thompson NFA and simulate it against each input word, or
//   - precompile once with automata.NFAToDFA (and perhaps automata.Minimize),
//     then simulate a DFA per input.
//
// There is no error return because every Node built from the six constructors
// denotes a regular language, and both strategies can build a matcher for it.
func FromAST(n regex.Node) *Regexp {
	panic("unimplemented: phase 1.9 FromAST")
}

// Contains reports whether the entire word w is in the language of r,
// satisfying automata.Language.
//
// This is whole-word acceptance, not search: read it as "the language contains
// w", not "w contains a match", which is what Go's regexp.MatchString asks.
//
// A Regexp has no alphabet: w may be any Word, and a word containing a symbol
// the pattern never mentions is rejected (Lit('a') does not accept "ba").
func (r *Regexp) Contains(w automata.Word) bool {
	panic("unimplemented: phase 1.9 Regexp.Contains")
}
