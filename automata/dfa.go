package automata

// Phase 1.1 — the deterministic finite automaton: the 5-tuple (Q, Sigma, delta, q0,
// F) of Sipser Definition 1.5.

// DFA is a deterministic finite automaton.
//
// The contract fixes nothing about this struct. Represent Q, delta, and F
// however you like; NewDFA translates the 5-tuple into your representation.
//
// At phase 1.7, regex.FromDFA has to read a DFA from outside this package, so plan
// now to expose its states, transitions, start state, and accepting states.
// Prefer read-only methods to exported fields; the README's "Design
// considerations" explains why.
type DFA struct {
	// your representation here
}

// NewDFA builds a DFA from an explicit 5-tuple, validating it.
//
// It returns an error for:
//   - an alphabet member that is not an ASCII letter or digit
//   - a state listed twice in states (a repeat in alphabet or accepting is
//     fine: only states defines labels)
//   - a start state not in states
//   - an accepting state not in states
//   - a transition whose From is not in states, or whose On is not in alphabet
//   - a transition whose target is not in states
//   - a delta that is not total: every (state, symbol) pair in
//     states x alphabet needs an image
//
// It must accept an empty alphabet, unlike Sipser; delta is then empty (or
// nil) and totality holds trivially.
func NewDFA(states []State, alphabet []Symbol, delta map[Transition]State,
	start State, accepting []State) (*DFA, error) {
	panic("unimplemented: phase 1.1 NewDFA")
}

// Run reports whether d accepts all of w.
//
// A word containing a symbol outside d's alphabet is not in Sigma*, so Run
// rejects it.
func (d *DFA) Run(w Word) bool {
	panic("unimplemented: phase 1.1 DFA.Run")
}

// Contains reports whether w is in L(d), satisfying Language. It agrees with
// Run on every word.
func (d *DFA) Contains(w Word) bool {
	panic("unimplemented: phase 1.1 DFA.Contains")
}

// NumStates returns |Q|: the number of states given to NewDFA, including any
// unreachable and dead states.
func (d *DFA) NumStates() int {
	panic("unimplemented: phase 1.1 DFA.NumStates")
}

// Alphabet returns d's alphabet as a set: each symbol exactly once, in any
// order. The slice is a fresh copy that belongs to the caller; modifying it
// must not change d. An empty alphabet gives a zero-length (or nil) slice.
//
// Use it to check that a construction preserved its input's alphabet, and to
// compile an expression over the same alphabet as a machine you already have:
//
//	Compile(FromDFA(d), d.Alphabet())
func (d *DFA) Alphabet() []Symbol {
	panic("unimplemented: phase 1.1 DFA.Alphabet")
}
