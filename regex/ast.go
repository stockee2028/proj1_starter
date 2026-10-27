package regex

import "strings"

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

type emptyset struct{}

func (*emptyset) isNode() {}

type eps struct{}

func (*eps) isNode() {}

type lit struct{ c byte }

func (*lit) isNode() {}

type seq struct{ ns []Node }

func (*seq) isNode() {}

type or struct{ ns []Node }

func (*or) isNode() {}

type star struct{ n Node }

func (*star) isNode() {}

const (
	precOr     = 1
	precConcat = 2
	precStar   = 3
)

func nodePrec(n Node) int {
	switch n.(type) {
	case *or:
		return precOr
	case *seq:
		return precConcat
	case *star:
		return precStar
	default:
		return precStar + 1
	}
}

func render(n Node, parentPrec int) string {
	if n == nil {
		return ""
	}

	var s string
	switch x := n.(type) {
	case *emptyset:
		s = "∅"
	case *eps:
		s = "ε"
	case *lit:
		s = string(x.c)
	case *seq:
		if len(x.ns) == 0 {
			s = "ε"
		} else {
			var b strings.Builder
			for _, child := range x.ns {
				b.WriteString(render(child, precConcat))
			}
			s = b.String()
		}
	case *or:
		if len(x.ns) == 0 {
			s = "∅"
		} else {
			var b strings.Builder
			for i, child := range x.ns {
				if i > 0 {
					b.WriteString("|")
				}
				b.WriteString(render(child, precOr))
			}
			s = b.String()
		}
	case *star:
		child := render(x.n, precStar)
		if _, ok := x.n.(*star); ok {
			child = "(" + child + ")"
		}
		s = child + "*"
	default:
		panic("unknown regex node")
	}

	if nodePrec(n) < parentPrec {
		return "(" + s + ")"
	}
	return s
}

// Emptyset returns the regular expression matching no words at all.
//
// Note that this is not the same as Eps(): the empty set contains no words, and
// epsilon's language contains exactly one, the empty word.
func Emptyset() Node {
	return &emptyset{}
}

// Eps returns the regular expression matching exactly the empty word.
func Eps() Node {
	return &eps{}
}

// Lit returns the regular expression matching exactly the one-symbol word c.
// It panics unless c is an ASCII letter or digit (A-Z, a-z, or 0-9).
func Lit(c byte) Node {
	if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')) {
		panic("regex literal must be ASCII letter or digit")
	}
	return &lit{c: c}
}

// Seq returns the concatenation of ns, left to right. Seq() with no arguments
// denotes {ε}, like Eps(), and renders as ε: concatenating nothing leaves the
// empty word.
func Seq(ns ...Node) Node {
	return &seq{ns: ns}
}

// Or returns the union of ns. Or() with no arguments denotes the empty
// language, like Emptyset(), and renders as ∅: the union of nothing is empty.
func Or(ns ...Node) Node {
	return &or{ns: ns}
}

// Star returns the Kleene star of n.
func Star(n Node) Node {
	return &star{n: n}
}

func (s *seq) String() string {
	return render(s, 0)
}

func (o *or) String() string {
	return render(o, 0)
}

func (s *star) String() string {
	return render(s, 0)
}

func (e *emptyset) String() string { return render(e, 0) }

func (e *eps) String() string { return render(e, 0) }

func (l *lit) String() string { return render(l, 0) }
