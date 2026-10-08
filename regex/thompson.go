package regex

import "401/automata"

// Phase 1.6 — Thompson's construction (Sipser Lemma 1.55): one NFA fragment per AST
// node type.

// Compile returns an NFA over exactly the alphabet sigma that accepts the
// language of n, by Thompson's construction.
//
// The caller supplies sigma because an expression names a language but not an
// alphabet. Every symbol occurring in n must be in sigma. sigma may contain
// more (those symbols simply have no transitions), and it may be empty or nil
// if n mentions no symbol: Compile(Eps(), nil) is legal. Compile panics if
// sigma contains a byte that is not an ASCII letter or digit, or omits a
// literal occurring in n. It leaves n unchanged, and n may share subtrees.
func Compile(n Node, sigma []automata.Symbol) *automata.NFA {
	panic("unimplemented: phase 1.6 Compile")
}
