package oracle

import (
	"regexp"
	"testing"

	"401/automata"
	"401/regex"
)

// CheckAgainstStdlib compares m against Go's regexp package on each word,
// treating the standard library as ground truth for the language of n.
//
// m is an automata.Language, so the same harness checks the NFA your phase 1.6
// Compile builds as well as the engine.Regexp of phase 1.9 — and it compiles and is
// usable before phase 1.9 exists.
//
// Two details here are the difference between a working oracle and an
// afternoon of confusion:
//
// The pattern is anchored. regexp.MatchString reports whether the input
// *contains* a match, while Contains asks whether the language contains the
// whole word. Comparing unanchored, every non-trivial pattern disagrees
// immediately.
//
// The AST is the input, not a pattern string. The caller builds a regex.Node,
// this function renders it with String(), translates ε and ∅ with GoPattern,
// and hands the result to regexp while the caller hands the same Node to the
// engine. Literals are ASCII letters/digits, so no escaping or input encoding
// is needed. Emptyset is supported, including inside larger expressions.
func CheckAgainstStdlib(t *testing.T, m automata.Language, n regex.Node, words []automata.Word) {
	t.Helper()

	pattern := n.String()
	ref, err := regexp.Compile(`\A(?:` + GoPattern(pattern) + `)\z`)
	if err != nil {
		t.Fatalf("String() produced %q, which is invalid after translating course syntax: %v\n"+
			"this is a bug in your String(), not in the harness", pattern, err)
	}

	for _, w := range words {
		got := m.Contains(w)
		want := ref.MatchString(string(w))
		if got != want {
			t.Errorf("pattern %q on %q: Contains said %v, stdlib said %v", pattern, string(w), got, want)
		}
	}
}
