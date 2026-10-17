package automata

// Phase 1.2 — closure operations, and the decision procedures built from them: the
// product construction (Sipser Theorem 1.25) and E_DFA and EQ_DFA (Sipser
// Theorems 4.4 and 4.5).

// Intersect returns a DFA for L(a) intersect L(b), by the product construction.
//
// a and b must have the same alphabet, and the result has that alphabet too.
// Alphabets are compared as sets: neither the order of the symbols passed to
// NewDFA nor repeats among them matter, so machines built over {0,1} and
// {1,0,1} have the same alphabet. If the alphabets differ as sets, even when
// one contains the other, Intersect returns an error.
func Intersect(a, b *DFA) (*DFA, error) {
	if a == nil || b == nil {
		return nil, err("nil DFA")
	}

	if !sameAlphabet(a.Alphabet(), b.Alphabet()) {
		return nil, err("alphabets differ")
	}

	sigma := a.Alphabet()
	if len(sigma) == 0 {
		// Over the empty alphabet, the only possible word is ε.
		if _, okA := a.accepting[a.start]; okA {

			if _, okB := b.accepting[b.start]; okB {
				return NewDFA([]State{0}, nil, map[Transition]State{}, 0, []State{0})
			}
		}

		return NewDFA([]State{0}, nil, map[Transition]State{}, 0, nil)
	}

	type pair struct {
		q1 State
		q2 State
	}

	start := pair{a.start, b.start}
	queue := []pair{start}
	seen := map[pair]State{start: 0}
	states := []State{0}
	delta := make(map[Transition]State)
	accepting := []State{}
	acceptingSet := make(map[State]struct{})

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		curr := seen[p]

		_, aOK := a.accepting[p.q1]
		_, bOK := b.accepting[p.q2]

		if aOK && bOK {
			acceptingSet[curr] = struct{}{}
		}

		for _, sym := range sigma {

			next1 := a.delta[Transition{From: p.q1, On: sym}]
			next2 := b.delta[Transition{From: p.q2, On: sym}]
			next := pair{next1, next2}

			if _, ok := seen[next]; !ok {
				seen[next] = State(len(states))
				states = append(states, State(len(states)))
				queue = append(queue, next)
			}

			delta[Transition{From: curr, On: sym}] = seen[next]
		}
	}

	for s := range acceptingSet {
		accepting = append(accepting, s)
	}

	return NewDFA(states, sigma, delta, 0, accepting)
}

func sameAlphabet(a, b []Symbol) bool {
	if len(a) != len(b) {
		return false
	}

	seen := make(map[Symbol]int, len(a))

	for _, s := range a {
		seen[s]++
	}

	for _, s := range b {
		if seen[s] == 0 {
			return false
		}
		seen[s]--
	}

	return true
}

// Union returns a DFA for L(a) union L(b).
//
// Alphabets work as in Intersect: a and b must have the same alphabet,
// compared as sets, and Union returns an error if they don't.
func Union(a, b *DFA) (*DFA, error) {
	if a == nil || b == nil {
		return nil, err("nil DFA")
	}

	if !sameAlphabet(a.Alphabet(), b.Alphabet()) {
		return nil, err("alphabets differ")
	}

	sigma := a.Alphabet()
	if len(sigma) == 0 {

		if _, okA := a.accepting[a.start]; okA {
			return NewDFA([]State{0}, nil, map[Transition]State{}, 0, []State{0})
		}

		if _, okB := b.accepting[b.start]; okB {
			return NewDFA([]State{0}, nil, map[Transition]State{}, 0, []State{0})
		}

		return NewDFA([]State{0}, nil, map[Transition]State{}, 0, nil)
	}

	type pair struct {
		q1 State
		q2 State
	}

	start := pair{a.start, b.start}
	queue := []pair{start}
	seen := map[pair]State{start: 0}
	states := []State{0}
	delta := make(map[Transition]State)
	acceptingSet := make(map[State]struct{})

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		curr := seen[p]

		_, aOK := a.accepting[p.q1]
		_, bOK := b.accepting[p.q2]

		if aOK || bOK {
			acceptingSet[curr] = struct{}{}
		}

		for _, sym := range sigma {
			next1 := a.delta[Transition{From: p.q1, On: sym}]
			next2 := b.delta[Transition{From: p.q2, On: sym}]
			next := pair{next1, next2}

			if _, ok := seen[next]; !ok {
				seen[next] = State(len(states))
				states = append(states, State(len(states)))
				queue = append(queue, next)
			}

			delta[Transition{From: curr, On: sym}] = seen[next]
		}
	}

	accepting := make([]State, 0, len(acceptingSet))

	for s := range acceptingSet {
		accepting = append(accepting, s)
	}

	return NewDFA(states, sigma, delta, 0, accepting)
}

// Complement returns a DFA for the complement of L(a), over a's alphabet.
//
// The complement is relative to Sigma*, where Sigma is a's alphabet: a word
// containing a symbol outside Sigma is in neither L(a) nor its complement.
func Complement(a *DFA) *DFA {

	if a == nil {
		return &DFA{

			states:    []State{0},
			alphabet:  nil,
			delta:     map[Transition]State{},
			start:     0,
			accepting: map[State]struct{}{0: {}},
		}
	}

	accepting := make([]State, 0, len(a.states))
	for _, s := range a.states {
		if _, ok := a.accepting[s]; !ok {
			accepting = append(accepting, s)
		}
	}

	c, err := NewDFA(

		append([]State(nil), a.states...),
		append([]Symbol(nil), a.alphabet...),
		copyDeltaMap(a.delta),
		a.start,
		accepting,
	)
	if err != nil {
		panic(err)
	}
	return c
}

// Empty reports whether L(a) is the empty language.
func Empty(a *DFA) bool {

	if a == nil {
		return true
	}

	seen := map[State]struct{}{a.start: {}}
	queue := []State{a.start}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		if _, ok := a.accepting[q]; ok {
			return false
		}

		for _, sym := range a.alphabet {
			next := a.delta[Transition{From: q, On: sym}]
			if _, ok := seen[next]; ok {
				continue
			}
			seen[next] = struct{}{}
			queue = append(queue, next)
		}
	}
	return true
}

// Equivalent reports whether L(a) equals L(b), by deciding whether their
// symmetric difference is empty:
//
//	Equivalent(a, b)  <=>  Empty(Intersect(a, Complement(b)))
//	                  and  Empty(Intersect(Complement(a), b))
//
// It panics if a and b have different alphabets, since it has no error return
// to report one with. It returns false exactly when Witness(a, b) reports a
// word.
func Equivalent(a, b *DFA) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	if !sameAlphabet(a.Alphabet(), b.Alphabet()) {
		panic("alphabets differ")
	}

	_, ok := Witness(a, b)
	return !ok
}

// Witness returns a shortest word on which a and b disagree, and true, or the
// zero Word and false if L(a) equals L(b).
//
// Shortest means by length; when several words tie, any of them will do. The
// empty word can be a witness, so check the bool, not the word.
//
// Witness panics if a and b have different alphabets. It reports a word exactly
// when Equivalent(a, b) is false.
func Witness(a, b *DFA) (Word, bool) {

	if a == nil || b == nil {
		if a == nil && b == nil {
			return nil, false
		}
		return nil, true
	}

	if !sameAlphabet(a.Alphabet(), b.Alphabet()) {
		panic("alphabets differ")
	}

	type pair struct {
		q1 State
		q2 State
	}

	start := pair{a.start, b.start}
	queue := []pair{start}
	wordByState := map[pair]Word{start: nil}
	seen := map[pair]struct{}{start: {}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		w := wordByState[p]

		aOK := false

		if _, ok := a.accepting[p.q1]; ok {
			aOK = true
		}
		bOK := false
		if _, ok := b.accepting[p.q2]; ok {
			bOK = true
		}
		if aOK != bOK {
			return w, true
		}

		for _, sym := range a.alphabet {
			next := pair{
				a.delta[Transition{From: p.q1, On: sym}],
				b.delta[Transition{From: p.q2, On: sym}],
			}

			if _, ok := seen[next]; ok {
				continue
			}
			
			seen[next] = struct{}{}
			wordByState[next] = append(append(Word(nil), w...), sym)
			queue = append(queue, next)
		}
	}

	return nil, false
}
