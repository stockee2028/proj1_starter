// Package regex is the regular-expression layer of the Project 1 toolkit: the
// abstract syntax tree and its rendering (phase 1.5), Thompson's construction into
// an NFA (phase 1.6), and state elimination back out to an expression (phase 1.7).
//
// There is no parser in Project 1. A pattern is a Go expression built from
// the six AST constructors:
//
//	// (a|b)*abb
//	Seq(Star(Or(Lit('a'), Lit('b'))), Lit('a'), Lit('b'), Lit('b'))
//
// Project 2 adds Parse to this package; nothing here changes when it does.
//
// The conventions that apply across the whole contract, including the
// rendering rules for String, are in the repository README.
package regex
