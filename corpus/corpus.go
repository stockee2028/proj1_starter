package corpus

import "401/automata"

// Sigma is the alphabet shared by every machine in this package.
var Sigma = []automata.Symbol{'0', '1'}

// Machine is a named corpus DFA together with the facts a test can assert
// about it.
type Machine struct {
	// Name is stable across offerings; tests may refer to it.
	Name string

	// Desc states the language in words.
	Desc string

	// New builds a fresh DFA. Always a fresh one: a test that mutates a shared
	// machine would poison every later test in the run.
	New func() *automata.DFA

	// MinStates is the Myhill-Nerode class count — the state count of the
	// minimal DFA for this language.
	MinStates int
}

// All returns every machine in the corpus, in a stable order.
func All() []Machine {
	return []Machine{
		{
			Name:      "EvenZeros",
			Desc:      "an even number of 0s",
			New:       EvenZeros,
			MinStates: 2,
		},
		{
			Name:      "EndsWith01",
			Desc:      "ends in 01",
			New:       EndsWith01,
			MinStates: 3,
		},
		{
			Name:      "DivisibleBy3",
			Desc:      "binary numerals divisible by 3, most significant bit first",
			New:       DivisibleBy3,
			MinStates: 3,
		},
		{
			Name:      "Contains001",
			Desc:      "contains 001 as a substring",
			New:       Contains001,
			MinStates: 4,
		},
		{
			Name:      "CanonicalBinary",
			Desc:      "binary numerals with no leading 0: the word 0, or a 1 followed by anything",
			New:       CanonicalBinary,
			MinStates: 4,
		},
		{
			Name:      "DivisibleBy6",
			Desc:      "binary numerals divisible by 6, most significant bit first, over one state per residue",
			New:       DivisibleBy6,
			MinStates: 4,
		},
		{
			Name:      "ThirdFromLast1",
			Desc:      "the third symbol from the end is 1",
			New:       ThirdFromLast1,
			MinStates: 8,
		},
	}
}

// EvenZeros accepts words containing an even number of 0s.
//
// Two states, tracking the parity of the 0-count. The empty word is accepted:
// zero is even.
func EvenZeros() *automata.DFA {
	const (
		even automata.State = 0
		odd  automata.State = 1
	)
	return mustDFA(automata.NewDFA(
		[]automata.State{even, odd},
		Sigma,
		map[automata.Transition]automata.State{
			{From: even, On: '0'}: odd,
			{From: even, On: '1'}: even,
			{From: odd, On: '0'}:  even,
			{From: odd, On: '1'}:  odd,
		},
		even,
		[]automata.State{even},
	))
}

// EndsWith01 accepts words ending in the two-symbol suffix 01.
//
// Three states: no useful suffix seen, last symbol was 0, last two were 01.
func EndsWith01() *automata.DFA {
	const (
		none       automata.State = 0 // last symbol was 1, or no symbols yet
		gotZero    automata.State = 1 // last symbol was 0
		gotZeroOne automata.State = 2 // last two symbols were 01
	)
	return mustDFA(automata.NewDFA(
		[]automata.State{none, gotZero, gotZeroOne},
		Sigma,
		map[automata.Transition]automata.State{
			{From: none, On: '0'}:       gotZero,
			{From: none, On: '1'}:       none,
			{From: gotZero, On: '0'}:    gotZero,
			{From: gotZero, On: '1'}:    gotZeroOne,
			{From: gotZeroOne, On: '0'}: gotZero,
			{From: gotZeroOne, On: '1'}: none,
		},
		none,
		[]automata.State{gotZeroOne},
	))
}

// DivisibleBy3 accepts binary numerals divisible by 3, read most significant
// bit first.
//
// Three states, one per residue mod 3, with delta(r, b) = (2r + b) mod 3:
// appending a bit doubles the value and adds the bit.
//
// The empty word is accepted, on the convention that it denotes 0. That is a
// convention rather than a theorem — a machine that rejects the empty word here
// would need four states and is a different language.
func DivisibleBy3() *automata.DFA {
	const (
		r0 automata.State = 0
		r1 automata.State = 1
		r2 automata.State = 2
	)
	return mustDFA(automata.NewDFA(
		[]automata.State{r0, r1, r2},
		Sigma,
		map[automata.Transition]automata.State{
			{From: r0, On: '0'}: r0, // 0 -> 0
			{From: r0, On: '1'}: r1, // 0 -> 1
			{From: r1, On: '0'}: r2, // 1 -> 2
			{From: r1, On: '1'}: r0, // 1 -> 3 = 0
			{From: r2, On: '0'}: r1, // 2 -> 4 = 1
			{From: r2, On: '1'}: r2, // 2 -> 5 = 2
		},
		r0,
		[]automata.State{r0},
	))
}

// Contains001 accepts words containing 001 as a substring. It is Sipser's E2
// (Example 1.21), state for state.
//
// Four states, recording how much of the pattern was just seen: nothing, 0, 00,
// and the whole of 001 — an accepting sink, since once the pattern has appeared
// no later input can undo it. Note the 0 out of q00: after 000 the last two
// symbols are still 00, so the machine stays put rather than starting over.
func Contains001() *automata.DFA {
	const (
		q    automata.State = 0 // no part of the pattern just seen
		q0   automata.State = 1 // just seen 0
		q00  automata.State = 2 // just seen 00
		q001 automata.State = 3 // seen 001 somewhere
	)
	return mustDFA(automata.NewDFA(
		[]automata.State{q, q0, q00, q001},
		Sigma,
		map[automata.Transition]automata.State{
			{From: q, On: '0'}:    q0,
			{From: q, On: '1'}:    q,
			{From: q0, On: '0'}:   q00,
			{From: q0, On: '1'}:   q,
			{From: q00, On: '0'}:  q00,
			{From: q00, On: '1'}:  q001,
			{From: q001, On: '0'}: q001,
			{From: q001, On: '1'}: q001,
		},
		q,
		[]automata.State{q001},
	))
}

// CanonicalBinary accepts binary numerals written without leading zeros: the
// word 0, or a 1 followed by anything.
//
// Four states, one of them dead: a numeral that starts with 0 has to stop
// there, so any symbol after a leading 0 is unrecoverable. The empty word is
// rejected — it is not a numeral — and so is 00.
//
// Intersected with DivisibleBy3 it gives the multiples of 3 as they are
// normally written, without DivisibleBy3's convention about the empty word.
func CanonicalBinary() *automata.DFA {
	const (
		start  automata.State = 0
		zero   automata.State = 1 // read exactly 0
		number automata.State = 2 // read a 1 and then anything
		dead   automata.State = 3 // read a 0 and then something more
	)
	return mustDFA(automata.NewDFA(
		[]automata.State{start, zero, number, dead},
		Sigma,
		map[automata.Transition]automata.State{
			{From: start, On: '0'}:  zero,
			{From: start, On: '1'}:  number,
			{From: zero, On: '0'}:   dead,
			{From: zero, On: '1'}:   dead,
			{From: number, On: '0'}: number,
			{From: number, On: '1'}: number,
			{From: dead, On: '0'}:   dead,
			{From: dead, On: '1'}:   dead,
		},
		start,
		[]automata.State{zero, number},
	))
}

// DivisibleBy6 accepts binary numerals divisible by 6, read most significant
// bit first, with the empty word denoting 0 as in DivisibleBy3.
//
// Six states, one per residue mod 6, with delta(r, b) = (2r + b) mod 6: the
// construction from DivisibleBy3, unchanged but for the modulus. Unlike every
// other machine in this package, it is not minimal. MinStates is 4, so
// two pairs of residues must be indistinguishable by any suffix. Which pairs,
// and why, is a Myhill-Nerode question worth answering before you trust your
// Minimize on it.
func DivisibleBy6() *automata.DFA {
	const n = 6
	states := make([]automata.State, n)
	delta := make(map[automata.Transition]automata.State)
	for r := range n {
		states[r] = automata.State(r)
		for _, c := range Sigma {
			b := int(c - '0')
			delta[automata.Transition{From: automata.State(r), On: c}] = automata.State((2*r + b) % n)
		}
	}
	return mustDFA(automata.NewDFA(states, Sigma, delta, 0, []automata.State{0}))
}

// ThirdFromLast1 accepts words whose third symbol from the end is 1. It is the
// language of Sipser's Example 1.30.
//
// Eight states, one for each possible last three symbols, read as a 3-bit
// number; a word shorter than three is padded on the left with 0s, which is
// right, since it has no third-from-last symbol to be 1. Reading bit b takes
// state s to (2s + b) mod 8, and the accepting states are those whose high bit
// is set.
//
// An NFA for this language needs only four states: it guesses which 1 is the
// third from last and checks that exactly two symbols follow. MinStates is 8,
// so no DFA does better than this one, and the subset construction on that
// four-state NFA has to produce at least eight. That gap is the exponential
// blowup of determinization at the smallest size where it is visible.
func ThirdFromLast1() *automata.DFA {
	const n = 8
	states := make([]automata.State, n)
	delta := make(map[automata.Transition]automata.State)
	var accepting []automata.State
	for s := range n {
		states[s] = automata.State(s)
		if s >= n/2 {
			accepting = append(accepting, automata.State(s))
		}
		for _, c := range Sigma {
			b := int(c - '0')
			delta[automata.Transition{From: automata.State(s), On: c}] = automata.State((2*s + b) % n)
		}
	}
	return mustDFA(automata.NewDFA(states, Sigma, delta, 0, accepting))
}

// mustDFA panics if the corpus itself is malformed.
//
// A panic here means one of two things: either phase 1.1 NewDFA is not implemented
// yet (expected early on), or NewDFA is rejecting a 5-tuple that is in fact
// valid — which is worth investigating, since these machines are correct.
func mustDFA(d *automata.DFA, err error) *automata.DFA {
	if err != nil {
		panic("corpus: NewDFA rejected a valid machine: " + err.Error())
	}
	return d
}
