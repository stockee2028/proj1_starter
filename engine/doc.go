// Package engine is the regular-expression matcher (phase 1.9): it turns an AST
// into a reusable Regexp that decides whole-word membership.
//
// Project 2 adds Compile(pattern string) (*Regexp, error), which composes a
// parser with FromAST. In Project 1 a pattern is a Go expression.
//
// The conventions that apply across the whole contract are in the repository
// README.
package engine
