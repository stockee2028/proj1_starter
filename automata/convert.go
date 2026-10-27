package automata

import (
	"fmt"
	"sort"
	"strings"
)

// Phases 1.3 and 1.4 — the conversions between NFAs and DFAs.

// DFAToNFA returns an NFA with the same language and alphabet as d.
// It leaves d unchanged.
func DFAToNFA(d *DFA) *NFA {

	alpha := append([]Symbol(nil), d.alphabet...)
	states := append([]State(nil), d.states...)
	delta := make(map[Transition][]State, len(d.delta))

	for tr, target := range d.delta {
		delta[tr] = []State{target}
	}

	accepting := d.Accepting()
	n, err := NewNFA(states, alpha, delta, nil, d.start, accepting)

	if err != nil {
		panic(err)
	}

	return n
}

// NFAToDFA returns a DFA with the same language and alphabet as n, by the
// subset construction (Sipser Theorem 1.39). It leaves n unchanged.
//
// It performs the construction lazily, building only reachable DFA states.
//
// The result has at most one state per subset of n's states reachable from the
// start subset, plus possibly the empty subset, and the same NumStates() on
// every call with the same n.  nextSet
func NFAToDFA(n *NFA) *DFA {

	sigma := append([]Symbol(nil), n.alphabet...)

	startSet := n.epsilonClosure(map[State]struct{}{n.start: {}})
	seen := map[string]State{stateSetKey(startSet): 0}
	stateToSet := map[State]map[State]struct{}{0: startSet}
	queue := []State{0}
	states := []State{0}
	delta := make(map[Transition]State)

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		currSet := stateToSet[q]

		for _, a := range sigma {
			nextSet := make(map[State]struct{})

			for s := range currSet {
				for _, target := range n.delta[Transition{From: s, On: a}] {
					for t := range n.epsilonClosure(map[State]struct{}{target: {}}) {
						nextSet[t] = struct{}{}
					}
				}
			}

			key := stateSetKey(nextSet)
			next, ok := seen[key]

			if !ok {
				next = State(len(states))
				seen[key] = next
				stateToSet[next] = nextSet
				states = append(states, next)
				queue = append(queue, next)
			}

			delta[Transition{From: q, On: a}] = next
		}
	}

	accepting := make([]State, 0, len(states))


	for q, set := range stateToSet {
		for s := range set {

			if _, ok := n.accepting[s]; ok {
				accepting = append(accepting, q)
				break
			}
		}
	}

	if len(sigma) == 0 {
		// Over the empty alphabet, the DFA has a single state that represents the
		// epsilon-closure of the start state. There are no symbol transitions to
		// add, and totality is trivial.
		if _, ok := n.accepting[n.start]; ok {
			if d, err := NewDFA([]State{0}, nil, map[Transition]State{}, 0, []State{0}); err == nil {
				return d
			}
		}

		if d, err := NewDFA([]State{0}, nil, map[Transition]State{}, 0, nil); err == nil {
			return d
		}

		panic("failed to build empty-alphabet DFA")
	}

	if len(states) == 0 {
		states = []State{0}
	}

	d, err := NewDFA(states, sigma, delta, 0, accepting)

	if err != nil {
		panic(err)
	}

	return d
}

func equalStateSets(a, b map[State]struct{}) bool {

	if len(a) != len(b) {
		return false
	}

	for s := range a {
		if _, ok := b[s]; !ok {
			return false
		}
	}

	return true
}

func stateSetKey(set map[State]struct{}) string {
	if len(set) == 0 {
		return "{}"
	}

	keys := make([]int, 0, len(set))
	
	for s := range set {
		keys = append(keys, int(s))
	}

	sort.Ints(keys)
	var b strings.Builder
	b.WriteString("{")

	for i, k := range keys {
		if i > 0 {
			b.WriteString(",")
		}

		b.WriteString(fmt.Sprintf("%d", k))
	}

	b.WriteString("}")


	return b.String()
}
