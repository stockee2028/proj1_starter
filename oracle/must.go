package oracle

// Must drops an error return, panicking if there is one.
//
// It exists so that closure operations compose in a single expression:
//
//	oracle.Must(automata.Intersect(
//	    corpus.EvenZeros(), automata.Complement(corpus.EndsWith01())))
//
// Without it, every such expression needs a temporary and an if err != nil,
// which in a test file is noise. Note that automata.Complement needs no Must:
// it cannot fail, because NewDFA already guarantees totality.
//
// Inside an oracle.Case body, a panic from Must becomes an ordinary test
// failure, so a malformed machine reports as a failed case rather than killing
// the run.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
