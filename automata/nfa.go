package automata

// Phase 1.3 — the nondeterministic finite automaton with epsilon-transitions
// (Sipser Definition 1.37).

// NFA is a nondeterministic finite automaton with epsilon-transitions.
//
// As with DFA, the contract fixes nothing about this struct.
type NFA struct {
	// your representation here
}

// NewNFA builds an NFA from an explicit 5-tuple plus an epsilon-transition map,
// validating it.
//
// It returns an error for:
//   - an alphabet member that is not an ASCII letter or digit
//   - a state listed twice in states (a repeat in alphabet, accepting, or a
//     target list is fine: only states defines labels)
//   - a start state not in states
//   - an accepting state not in states
//   - a transition whose From is not in states, or whose On is not in alphabet
//   - a transition or epsilon-transition whose source or any target is not in
//     states
//
// Unlike NewDFA there is no totality requirement: a missing key, a nil map, or
// an empty target list means no transition. Epsilon-cycles, self-loops, an
// empty accepting set, and an empty alphabet are all legal.
func NewNFA(states []State, alphabet []Symbol, delta map[Transition][]State,
	epsilon map[State][]State, start State, accepting []State) (*NFA, error) {
	panic("unimplemented: phase 1.3 NewNFA")
}

// Run reports whether n accepts all of w, by simulating the set of possible
// NFA computations in parallel.
//
// A word containing a symbol outside n's alphabet is not in Sigma*, so Run
// rejects it.
func (n *NFA) Run(w Word) bool {
	panic("unimplemented: phase 1.3 NFA.Run")
}

// Contains reports whether w is in L(n), satisfying Language. It agrees with
// Run on every word.
func (n *NFA) Contains(w Word) bool {
	panic("unimplemented: phase 1.3 NFA.Contains")
}

// NumStates returns |Q|: the number of states given to NewNFA, or built by the
// construction that returned n, unreachable states included.
func (n *NFA) NumStates() int {
	panic("unimplemented: phase 1.3 NFA.NumStates")
}
