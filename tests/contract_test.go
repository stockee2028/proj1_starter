package tests

import (
	"401/automata"
	"401/engine"
	"401/regex"
)

// Compile-time assertions that the contract still holds. These are compile-only
// checks, so if it compiles, then the contract is satisfied.

var (
	_ automata.Language = (*automata.FiniteLanguage)(nil)
	_ automata.Language = (*automata.DFA)(nil)
	_ automata.Language = (*automata.NFA)(nil)
	_ automata.Language = (*engine.Regexp)(nil)
)

// Methods that no interface above covers, asserted as method expressions.
var (
	_ func(*automata.DFA) int               = (*automata.DFA).NumStates
	_ func(*automata.DFA) []automata.Symbol = (*automata.DFA).Alphabet
	_ func(*automata.NFA) int               = (*automata.NFA).NumStates
)

// Free functions, asserted by taking their addresses at the declared types.
var (
	_ func(words ...automata.Word) *automata.FiniteLanguage = automata.NewFiniteLanguage

	_ func([]automata.State, []automata.Symbol, map[automata.Transition]automata.State,
		automata.State, []automata.State) (*automata.DFA, error) = automata.NewDFA
	_ func([]automata.State, []automata.Symbol, map[automata.Transition][]automata.State,
		map[automata.State][]automata.State, automata.State, []automata.State) (*automata.NFA, error) = automata.NewNFA

	_ func(a, b *automata.DFA) (*automata.DFA, error) = automata.Intersect
	_ func(a, b *automata.DFA) (*automata.DFA, error) = automata.Union
	_ func(a *automata.DFA) *automata.DFA             = automata.Complement
	_ func(a *automata.DFA) bool                      = automata.Empty
	_ func(a, b *automata.DFA) bool                   = automata.Equivalent
	_ func(a, b *automata.DFA) (automata.Word, bool)  = automata.Witness

	_ func(d *automata.DFA) *automata.NFA = automata.DFAToNFA
	_ func(n *automata.NFA) *automata.DFA = automata.NFAToDFA
	_ func(d *automata.DFA) *automata.DFA = automata.Minimize

	_ func() regex.Node                 = regex.Emptyset
	_ func() regex.Node                 = regex.Eps
	_ func(c byte) regex.Node           = regex.Lit
	_ func(ns ...regex.Node) regex.Node = regex.Seq
	_ func(ns ...regex.Node) regex.Node = regex.Or
	_ func(n regex.Node) regex.Node     = regex.Star

	_ func(n regex.Node, sigma []automata.Symbol) *automata.NFA = regex.Compile
	_ func(d *automata.DFA) regex.Node                          = regex.FromDFA
	_ func(n regex.Node) *engine.Regexp                         = engine.FromAST
)
