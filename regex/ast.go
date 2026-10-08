package regex

// Phase 1.5 — the regular-expression AST (Sipser Definition 1.52), including both
// the empty set and epsilon.

// Node is a regular expression, built only by the six constructors below.
//
// The unexported marker method keeps the set of node types closed: only types
// declared in this package can satisfy Node, so a tree walk over the six cases
// is exhaustive.
//
// String renders course regex syntax: ASCII letter/digit literals, ε, ∅,
// concatenation, |, *, and parentheses, with no escapes or whitespace. The
// README gives the rendering rules and precedence. String is graded, and the
// grading suites compare your later phases against Go's regexp through it.
// Parenthesize defensively: Seq(Or(Lit('a'), Lit('b')), Lit('c')) is (a|b)c,
// not a|bc, and extra parentheses are always allowed.
//
// A Node may appear more than once, in one tree or across several. Nothing
// modifies a Node, and String gives the same result on every call.
type Node interface {
	String() string
	isNode()
}

// Emptyset returns the regular expression matching no words at all.
//
// Note that this is not the same as Eps(): the empty set contains no words, and
// epsilon's language contains exactly one, the empty word.
func Emptyset() Node {
	panic("unimplemented: phase 1.5 Emptyset")
}

// Eps returns the regular expression matching exactly the empty word.
func Eps() Node {
	panic("unimplemented: phase 1.5 Eps")
}

// Lit returns the regular expression matching exactly the one-symbol word c.
// It panics unless c is an ASCII letter or digit (A-Z, a-z, or 0-9).
func Lit(c byte) Node {
	panic("unimplemented: phase 1.5 Lit")
}

// Seq returns the concatenation of ns, left to right. Seq() with no arguments
// denotes {ε}, like Eps(), and renders as ε: concatenating nothing leaves the
// empty word.
func Seq(ns ...Node) Node {
	panic("unimplemented: phase 1.5 Seq")
}

// Or returns the union of ns. Or() with no arguments denotes the empty
// language, like Emptyset(), and renders as ∅: the union of nothing is empty.
func Or(ns ...Node) Node {
	panic("unimplemented: phase 1.5 Or")
}

// Star returns the Kleene star of n.
func Star(n Node) Node {
	panic("unimplemented: phase 1.5 Star")
}
