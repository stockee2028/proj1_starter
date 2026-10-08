package oracle

import "401/regex"

// AST returns a random regex.Node of depth at most maxDepth.
//
// The generator's alphabet must contain only ASCII letters and digits. Keep
// it small — {'a','b'} is a useful choice — so subset constructions stay
// inspectable. Both epsilon and Emptyset are generated: GoPattern translates
// their course spellings for the stdlib oracle.
func (g *Gen) AST(maxDepth int) regex.Node {
	if maxDepth <= 1 {
		return g.leaf()
	}
	switch g.rng.IntN(6) {
	case 0, 1:
		return g.leaf()
	case 2:
		return regex.Seq(g.AST(maxDepth-1), g.AST(maxDepth-1))
	case 3:
		return regex.Or(g.AST(maxDepth-1), g.AST(maxDepth-1))
	case 4:
		return regex.Star(g.AST(maxDepth - 1))
	default:
		return regex.Seq(g.AST(maxDepth-1), g.AST(maxDepth-1), g.AST(maxDepth-1))
	}
}

// leaf returns a random literal, epsilon, or the empty set.
func (g *Gen) leaf() regex.Node {
	switch g.rng.IntN(8) {
	case 0:
		return regex.Eps()
	case 1:
		return regex.Emptyset()
	}
	return regex.Lit(g.alphabet[g.rng.IntN(len(g.alphabet))])
}
