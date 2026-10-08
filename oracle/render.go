package oracle

import "strings"

var goTokens = strings.NewReplacer("ε", "(?:)", "∅", `[^\s\S]`)

// GoPattern translates a course regex into an unanchored Go regexp.
//
// pattern must be valid course syntax: ASCII letter/digit literals, ε, ∅,
// concatenation, |, *, and parentheses. This is a token adapter, not a parser.
// Literal symbols cannot overlap syntax, so no escaping is needed. Epsilon
// becomes an empty group; the empty set becomes a character class excluding
// both whitespace and non-whitespace, which matches no character.
//
// Callers doing whole-word checks must anchor the result. CheckAgainstStdlib
// does this for student tests. Project 2 parses the original course string,
// not this translated Go pattern.
func GoPattern(pattern string) string {
	return goTokens.Replace(pattern)
}
