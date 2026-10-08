package automata

// Phases 1.3 and 1.4 — the conversions between NFAs and DFAs.

// DFAToNFA returns an NFA with the same language and alphabet as d.
// It leaves d unchanged.
func DFAToNFA(d *DFA) *NFA {
	panic("unimplemented: phase 1.3 DFAToNFA")
}

// NFAToDFA returns a DFA with the same language and alphabet as n, by the
// subset construction (Sipser Theorem 1.39). It leaves n unchanged.
//
// It performs the construction lazily, building only reachable DFA states.
//
// The result has at most one state per subset of n's states reachable from the
// start subset, plus possibly the empty subset, and the same NumStates() on
// every call with the same n.
func NFAToDFA(n *NFA) *DFA {
	panic("unimplemented: phase 1.4 NFAToDFA")
}
