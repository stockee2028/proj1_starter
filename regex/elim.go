package regex

import "401/automata"

// Phase 1.7 — state elimination (Sipser Lemma 1.60), via generalized NFAs whose
// transitions are labeled with regular expressions: rip out one state at a
// time, repairing the expressions on the paths through it, until only a start
// and an accept state remain.

// FromDFA returns a regular expression whose language is L(d), by state
// elimination. Any correct expression will do. It leaves d unchanged.
//
// FromDFA reads d from outside package automata, so it can use only what
// automata exports: NumStates, Alphabet, and whatever read access to the
// states, transitions, start state, and accepting states you added at phase 1.1.
//
// The resulting regular expression will not have symbols whose transitions all lead
// to a dead state: "no c" over {a,b,c} is (a|b)*. When you compile such an expression
// in tests, you'll still want to compile over d's alphabet so the resulting machines
// are comparable.
func FromDFA(d *automata.DFA) Node {
	panic("unimplemented: phase 1.7 FromDFA")
}
