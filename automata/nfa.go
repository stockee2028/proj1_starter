package automata

// Phase 1.3 — the nondeterministic finite automaton with epsilon-transitions
// (Sipser Definition 1.37).

// NFA is a nondeterministic finite automaton with epsilon-transitions.
//
// As with DFA, the contract fixes nothing about this struct.
type NFA struct {

	states	[]State
	alphabet	[]Symbol
	delta	map[Transition][]State
	epsilon	map[State][]State
	start	State
	accepting	map[State]struct{}
	
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

	seenStates := make(map[State]struct{}, len(states))

	for _, s := range states {
		if _, ok := seenStates[s]; ok {
			return nil, err("duplicate state in states")
		}

		seenStates[s] = struct{}{}
	}

	if _, ok := seenStates[start]; !ok {
		return nil, err("start state not in states")
	}

	alphaSet := make(map[Symbol]struct{}, len(alphabet))

	for _, a := range alphabet {
		if !isAsciiLetterDigit(a) {

			return nil, err("alphabet symbol not ASCII letter or digit")
		}

		alphaSet[a] = struct{}{}
	}

	acceptSet := make (map[State]struct{}, len(accepting))

	for _, s := range accepting {
		if _, ok := seenStates[s]; !ok {

			return nil, err("accepting state not in states")
		}
		acceptSet[s] = struct{}{}
	}

	for tr, targets := range delta {

		if _, ok := seenStates[tr.From]; !ok {
			return nil, err("transition From not in states")
		}

		if _, ok := alphaSet[tr.On]; !ok {
			return nil, err("transition On not in alphabet")
		}
		for _, target := range targets {
			if _, ok := seenStates[target]; !ok {
				return nil, err("transition target not in states")
			}
		}
	}

	for from, targets := range epsilon {

		if _, ok := seenStates[from]; !ok {
			return nil, err("epsilon-transition source not in states")
		}
		for _, target := range targets {

			if _, ok := seenStates[target]; !ok {

				return nil, err("epsilon-transition target not in states")
			}
		}
	}

	return &NFA{
		states:    append([]State(nil), states...),
		alphabet:  append([]Symbol(nil), alphabet...),
		delta:     copyDeltaSliceMap(delta),
		epsilon:   copyEpsilonMap(epsilon),
		start:     start,
		accepting: acceptSet,
	}, nil
}

func copyDeltaSliceMap(in map[Transition][]State) map[Transition][]State {

	if len(in) == 0 {
		return nil
	}

	out := make(map[Transition][]State, len(in))

	for k, v := range in {
		out[k] = append([]State(nil), v...)
	}


	return out
}

func copyEpsilonMap(in map[State][]State) map[State][]State {

	if len(in) == 0 {
		return nil
	}

	out := make(map[State][]State, len(in))

	for k, v := range in {
		out[k] = append([]State(nil), v...)
	}

	return out
}

func copyAcceptingSet(in []State) map[State]struct{} {
	out := make(map[State]struct{}, len(in))

	for _, s := range in {
		out[s] = struct{}{}
	}

	return out
}

func (n *NFA) epsilonClosure(states map[State]struct{}) map[State]struct{} {
	if n == nil {
		return nil
	}

	seen := make(map[State]struct{}, len(states))
	stack := make([]State, 0, len(states))

	for s := range states {
		seen[s] = struct{}{}
		stack = append(stack, s)
	}

	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, next := range n.epsilon[s] {

			if _, ok := seen[next]; ok {
				continue
			}

			seen[next] = struct{}{}
			stack = append(stack, next)
		}
	}

	return seen
}

// Run reports whether n accepts all of w, by simulating the set of possible
// NFA computations in parallel.
//
// A word containing a symbol outside n's alphabet is not in Sigma*, so Run
// rejects it.
func (n *NFA) Run(w Word) bool {
	if n == nil {
		return false
	}

	for _, sym := range w {
		if !containsSymbol(n.alphabet, sym) {
			return false
		}
	}

	current := map[State]struct{}{n.start: {}}
	current = n.epsilonClosure(current)

	for _, sym := range w {
		next := make(map[State]struct{})

		for state := range current {
			for _, target := range n.delta[Transition{From: state, On: sym}] {
				next[target] = struct{}{}
			}
		}
		if len(next) == 0 {
			return false

		}

		current = n.epsilonClosure(next)
	}

	for state := range current {

		if _, ok := n.accepting[state]; ok {
			return true
		}
	}

	return false
}

// Contains reports whether w is in L(n), satisfying Language. It agrees with
// Run on every word.
func (n *NFA) Contains(w Word) bool {
	return n.Run(w)
}

// NumStates returns |Q|: the number of states given to NewNFA, or built by the
// construction that returned n, unreachable states included.
func (n *NFA) NumStates() int {
	if n == nil {
		return 0
	}
	
	return len(n.states)
}
