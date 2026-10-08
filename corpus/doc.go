// Package corpus provides named DFAs, each with its known minimal state count.
// All lists every machine.
//
// Every machine is built by calling automata.NewDFA, so nothing here works
// until you have implemented phase 1.1: before then a corpus constructor panics
// with "unimplemented: phase 1.1 NewDFA", which is expected. A test that
// imports this package must live in tests/; the repository README explains
// why.
//
// MinStates on each machine is its Myhill-Nerode class count. Because the
// minimal DFA is unique, a Minimize result with a different count is not
// minimal.
package corpus
