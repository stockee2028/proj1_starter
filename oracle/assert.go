package oracle

import (
	"runtime/debug"
	"slices"
	"strings"
	"testing"

	"401/automata"
)

// Case runs body as a subtest, converting a panic into an ordinary test
// failure.
//
// A panic in a Go test normally ends the whole test binary, so every test
// after it never runs. The grading suites wrap every case in Case, so that an
// unimplemented stub or a crash in one function costs only the cases that
// reach it. The panic's stack trace is kept in the failure message.
//
// Case does not catch non-termination (go test -timeout ends the binary
// regardless), a fatal error such as a stack overflow, or a panic in a
// goroutine your code starts.
//
// Don't use it in your own tests. While you are developing, a panic is most
// useful delivered the way Go delivers it, with a stack trace and a stop.
func Case(t *testing.T, name string, body func(t *testing.T)) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		defer func() {
			v := recover()
			if v == nil {
				return
			}
			// An unimplemented stub needs no stack trace: the message already
			// names the function and the phase, and a wall of stacks makes an
			// early-in-the-project test run unreadable. A genuine crash is the
			// case where the stack is worth every line.
			if msg, ok := v.(string); ok && strings.HasPrefix(msg, "unimplemented:") {
				t.Errorf("%s", msg)
				return
			}
			t.Errorf("panicked: %v\n\n%s", v, debug.Stack())
		}()
		body(t)
	})
}

// Decides asserts that m's decision on w is want.
//
// m is any automata.Language (a FiniteLanguage, DFA, NFA, or engine.Regexp),
// and the decision goes through Contains. w is a plain Go string, converted to
// an automata.Word byte for byte.
func Decides(t *testing.T, m automata.Language, w string, want bool) {
	t.Helper()
	got := m.Contains(automata.Word(w))
	if got == want {
		return
	}
	verb := "reject"
	if want {
		verb = "accept"
	}
	t.Errorf("Contains(%q) = %v, want %v (expected the machine to %s)", w, got, want, verb)
}

// Accepts asserts that m accepts w.
func Accepts(t *testing.T, m automata.Language, w string) {
	t.Helper()
	Decides(t, m, w, true)
}

// Rejects asserts that m rejects w.
func Rejects(t *testing.T, m automata.Language, w string) {
	t.Helper()
	Decides(t, m, w, false)
}

// AssertEquivalent asserts that got and want accept the same language,
// reporting the shortest distinguishing word when they do not. If the two
// machines have different alphabets it reports that instead of letting
// Equivalent panic, since it usually means a construction dropped a symbol.
// It also reports Equivalent and Witness disagreeing with each other.
func AssertEquivalent(t *testing.T, got, want *automata.DFA) {
	t.Helper()
	// Equivalent requires both machines to be over the same alphabet and panics
	// otherwise, so check first: "a construction dropped a symbol" is a
	// diagnosis, where the panic is only a symptom.
	if g, w := symbolSet(got.Alphabet()), symbolSet(want.Alphabet()); g != w {
		t.Errorf("the two machines are not over the same alphabet (got %s, want %s), "+
			"so their languages cannot be compared. Every construction in the "+
			"contract preserves the alphabet of the machine it was given, and Compile "+
			"builds over the alphabet you hand it, so something upstream changed it.", g, w)
		return
	}
	if automata.Equivalent(got, want) {
		return
	}
	if w, ok := automata.Witness(got, want); ok {
		t.Errorf("languages differ; shortest distinguishing word is %q (got %v, want %v)",
			string(w), got.Contains(w), want.Contains(w))
		return
	}
	t.Errorf("Equivalent reported false but Witness found no distinguishing word — " +
		"one of the two is wrong")
}

// symbolSet renders an alphabet as a set, {0,1}, in byte order and without
// repeats, so that two alphabets equal as sets render identically.
func symbolSet(sigma []automata.Symbol) string {
	sorted := slices.Clone(sigma)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	parts := make([]string, len(sorted))
	for i, c := range sorted {
		parts[i] = string(c)
	}
	return "{" + strings.Join(parts, ",") + "}"
}
