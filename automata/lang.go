package automata

// This file is provided complete.
// The code you write for phase 1.0 is FiniteLanguage, in finitelang.go.

// Symbol is a single element of an alphabet. It is byte, not rune, so
// Word("0011") is exactly four symbols.
//
// DFA/NFA alphabets and regex literals are restricted to ASCII letters and
// digits (A-Z, a-z, 0-9); see the README.
type Symbol = byte

// Word is a finite sequence of symbols. It may hold any bytes, and a nil Word
// is the empty word.
type Word []Symbol

// State identifies a state of an automaton. Labels are arbitrary: they need
// not be contiguous, start at zero, or be non-negative.
type State int

// Transition is the domain of a transition function: a (state, symbol) pair.
type Transition struct {
	From State
	On   Symbol
}

// Language is anything that decides membership of a word. FiniteLanguage, DFA,
// NFA, and engine.Regexp implement it, and the oracle assertions accept any
// Language.
//
// Contains answers for every Word: a word that is not in the language,
// including one containing symbols the language never uses, gets false, never
// a panic.
type Language interface {
	Contains(Word) bool
}
