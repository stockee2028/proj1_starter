package automata

// Phase 1.1 — the deterministic finite automaton: the 5-tuple (Q, Sigma, delta, q0,
// F) of Sipser Definition 1.5.

// DFA is a deterministic finite automaton.
//
// The contract fixes nothing about this struct. Represent Q, delta, and F
// however you like; NewDFA translates the 5-tuple into your representation.
//
// At phase 1.7, regex.FromDFA has to read a DFA from outside this package, so plan
// now to expose its states, transitions, start state, and accepting states.// Prefer read-only methods to exported fields; the README's "Design// constransitionsiderations" explains why.
type DFA struct {
	states    []State
	alphabet  []Symbol
	delta     map[Transition]State
	start     State
	accepting map[State]struct{}
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

	for _, a := range alphabet {
		if !isAsciiLetterDigit(a) {
			return nil, err("alphabet symbol not ASCII letter or digit")
		}
	}

	acceptSet := make(map[State]struct{}, len(accepting))

	for _, s := range accepting {
		if _, ok := seenStates[s]; !ok {
			return nil, err("accepting state not in states")
		}
		acceptSet[s] = struct{}{}
	}

	if len(alphabet) == 0 {

		if len(delta) != 0 {

			for tr := range delta {

				if _, ok := seenStates[tr.From]; !ok {
					return nil, err("transition from is not in states")
				}

				if !isAsciiLetterDigit(tr.On) {
					return nil, err("transition symbol not in alphabet")
				}

				if _, ok := seenStates[delta[tr]]; !ok {
					return nil, err("transition target not in states")
				}
			}
		}

		return &DFA{
			states:    append([]State(nil), states...),
			alphabet:  append([]Symbol(nil), alphabet...),
			delta:     copyDeltaMap(delta),
			start:     start,
			accepting: acceptSet,
		}, nil
	}

	alphaSet := make(map[Symbol]struct{}, len(alphabet))

	for _, s := range alphabet {
		alphaSet[s] = struct{}{}
	}

	for tr, target := range delta {
		if _, ok := seenStates[tr.From]; !ok {
			return nil, err("transition from is not in states")
		}
		if _, ok := alphaSet[tr.On]; !ok {
			return nil, err("transition symbol not in alphabet")
		}
		if _, ok := seenStates[target]; !ok {
			return nil, err("transition target not in states")
		}
	}

	for _, q := range states {
		for _, a := range alphabet {
			if _, ok := delta[Transition{From: q, On: a}]; !ok {
				return nil, err("delta is not total")
			}
		}
	}

	return &DFA{

		states:    append([]State(nil), states...),
		alphabet:  append([]Symbol(nil), alphabet...),
		delta:     copyDeltaMap(delta),
		start:     start,
		accepting: acceptSet,
	}, nil
}

func err(msg string) error { return &dfaErr{msg: msg} }

type dfaErr struct{ msg string }

func (e *dfaErr) Error() string { return e.msg }

func isAsciiLetterDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}

	if b >= 'A' && b <= 'Z' {
		return true
	}

	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}

func copyDeltaMap(in map[Transition]State) map[Transition]State {
	if len(in) == 0 {
		return nil
	}
	out := make(map[Transition]State, len(in))

	for k, v := range in {
		out[k] = v
	}
	return out
}

// Run reports whether d accepts all of w.
//
// A word containing a symbol outside d's alphabet is not in Sigma*, so Run
// rejects it.
func (d *DFA) Run(w Word) bool {
	if d == nil {
		return false
	}

	state := d.start

	for _, sym := range w {
		if !containsSymbol(d.alphabet, sym) {
			return false
		}
		state = d.delta[Transition{From: state, On: sym}]
	}
	_, ok := d.accepting[state]

	return ok
}

func containsSymbol(alphabet []Symbol, s Symbol) bool {

	for _, sym := range alphabet {
		if sym == s {
			return true
		}
	}

	return false
}

// Contains reports whether w is in L(d), satisfying Language. It agrees with
// Run on every word.
func (d *DFA) Contains(w Word) bool {

	return d.Run(w)
}

// NumStates returns |Q|: the number of states given to NewDFA, including any
// unreachable and dead states.
func (d *DFA) NumStates() int {
	if d == nil {
		return 0
	}
	return len(d.states)
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

	if d == nil || len(d.alphabet) == 0 {
		if d == nil {
			return nil
		}
		return nil
	}

	out := make([]Symbol, len(d.alphabet))

	copy(out, d.alphabet)
	return out
}

// States returns the DFA's state labels in the order given to NewDFA.
func (d *DFA) States() []State {

	if d == nil {
		return nil
	}

	out := make([]State, len(d.states))
	copy(out, d.states)

	return out
}

// Start returns the start state.
func (d *DFA) Start() State {
	if d == nil {
		return 0
	}

	return d.start
}

// Accepting returns the accepting states as a fresh slice in the order first
// encountered in the constructor input, with duplicates removed.
func (d *DFA) Accepting() []State {
	if d == nil {
		return nil
	}

	out := make([]State, 0, len(d.accepting))
	seen := make(map[State]struct{}, len(d.accepting))

	for _, s := range d.states {
		if _, ok := d.accepting[s]; ok {
			if _, duplicate := seen[s]; duplicate {
				continue
			}
			seen[s] = struct{}{}
			out = append(out, s)
		}
	}
	return out;
}

// Step follows the DFA's transition on symbol a from state s.
func (d *DFA) Step(s State, a Symbol) State {

	if d == nil {
		return 0
	}

	if _, ok := d.delta[Transition{From: s, On: a}]; !ok {
		return 0
	}


	return d.delta[Transition{From: s, On: a}]
}
