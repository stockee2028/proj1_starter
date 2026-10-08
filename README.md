# Project 1: A Regular-Languages Toolkit

This repository holds Project 1: DFAs, NFAs, closure operations, both directions of Kleene's theorem, minimization, and a regular-expression engine. This README is the project's technical specification. The project handout on the course website covers the schedule, deliverables, reflections, grading, and the ideas behind the design.

The specification has two parts:

- **This README:** the repository rules, and every convention that the function signatures alone don't express.
- **The package documentation:** each contract function's doc comment gives its exact behavior. `go doc -all ./automata` shows a whole package, and `go doc ./automata NFAToDFA` shows one function. Do the same with `./regex` and `./engine` for the code you write, and `./oracle` and `./corpus` for the provided packages. Plain `go doc ./automata`, without `-all`, lists only the signatures.

The declarations that ship in the starter code for `automata`, `regex`, and `engine` are the contract. Match them exactly: spelling, receiver types, parameter and return types, error returns. `tests/contract_test.go` checks the signatures at compile time. Everything you add is yours to design: how you represent machines, your helper functions, and any extra exported methods you decide you need.

## Layout

```
401/
  go.mod      module 401
  automata/   DFA, NFA, closure operations, conversions, minimization   yours
  regex/      AST, Thompson's construction, state elimination           yours
  engine/     the AST-driven matcher                                    yours
  corpus/     named DFAs with known minimal state counts                provided
  oracle/     assertions, generators, reference harnesses               provided
  tests/      every test that imports corpus or oracle                  provided + yours
```

## Rules

**1. Leave these alone:**

- **`go.mod`.** The module path is `401`, so imports look like `import "401/automata"`. The grading suites are copied into your repo and import your packages by those paths.
- **`corpus/` and `oracle/`.** They contain no solution code. They're written against the contract, which is why the same packages work in everyone's repo.
- **`automata/lang.go`.** It ships complete: `Symbol`, `Word`, `State`, `Transition`, and `Language`.
- **The contract declarations in the stub files.** Names, receivers, parameters, and return values are fixed by the contract, while implementation is up to you.
- **`tests/contract_test.go`.** This checks that your code adheres to the contract, giving a compile error if it doesn't.

The grading suites assume all of these are exactly as shipped. Later projects add files to this repository assuming the same.

**2. Dependencies point one way: `engine` imports `regex`, and `regex` imports `automata`.** `automata` must never import `regex` or know anything about regex syntax. That separation is what lets a parser be added to `regex` later without touching `automata`.

**3. Any test that imports `corpus` or `oracle` goes in `tests/`.** It can't go in the project packages (e.g., `automata`) because `corpus` and `oracle` both import `automata`, and you'll get an import cycle.

In-package tests (`package automata` inside `automata/`) are fine and are the right tool for testing your unexported internals. Just keep them free of `corpus` and `oracle`, and build any machine they need right there with `NewDFA`. Declaring `package automata_test` instead also avoids the cycle, but a file like that can only see the exported API, which `tests/` already gives you.

## Workflow

```sh
go test ./...                        # every test in every package
go test ./tests/                     # one package
go test ./... -run 'TestP1_4'        # one phase, if you name tests by phase (below)
go test ./... -run 'TestP1_[0-4]'    # phases 1.0 through 1.4
go test ./... -failfast              # stop each package at its first failing test
go test ./... -run 'TestP1_4' -v     # add -v to list each test as it runs
```

Every contract function starts as a stub that panics with a message naming its phase, such as `unimplemented: phase 1.4 NFAToDFA`. That's the expected state for any phase you haven't built yet. Stubs panic rather than returning zero values so that "not written yet" is never mistaken for "written and wrong."

`oracle.Case` turns a panic into an ordinary test failure. The grading suites and `tests/smoke_test.go` wrap every case in it, so one unimplemented or crashing function doesn't stop the rest of the run. Don't use it in your own tests: while you're developing, a panic is most useful the way Go delivers it, with a stack trace and a stop.

Because your own tests don't use `oracle.Case`, a test that reaches a stub panics and ends its package's test run, and the tests after it never run. To test just what you're working on, use `-run`. It takes a regular expression and runs only the tests whose names match. The tests in `tests/smoke_test.go` are named by phase (`TestP1_0_Types`, `TestP1_1_Corpus`), and so are the grading suites. I recommend naming your own tests the same way so that `-run 'TestP1_4'` runs just your phase 1.4 tests.

`-failfast` stops at the first failing test instead of running the rest. It works one package at a time, so `go test ./... -failfast` will run tests for all packages and stops each package's tests at its own first failure.

The `Makefile` has shortcuts for a few common commands:

```sh
make smoke   # passes on a fresh clone, before you write anything
make test    # everything: go test ./...
make bench   # benchmarks only (phase 1.9)
make vet     # go vet ./...
```

## Contract conventions

These rules hold across the whole contract. Each function's doc comment adds the details specific to that function.

### Symbols, words, and alphabets

- **`Symbol` is `byte`, not `rune`.** `automata.Word("0011")` gives the four symbols you'd expect. AI may try to write you code that uses runes (Go's type that roughly corresponds to Unicode characters); make sure you're staying in byte-land.
- **Machine alphabets and regex literals use only ASCII letters and digits** (`A`-`Z`, `a`-`z`, `0`-`9`), and case matters. `NewDFA` and `NewNFA` return an error for any other alphabet member, even one no transition uses. `Lit` and `Compile` panic. `FiniteLanguage` stores explicit words and has no alphabet restriction.
- **The empty alphabet is legal** for `NewDFA`, `NewNFA`, and `Compile`, unlike in Sipser. Over the empty alphabet, $\Sigma^* = \{\varepsilon\}$, so a machine accepts either nothing or only the empty word, and DFA totality holds automatically. A pattern like `Star(Eps())` mentions no symbol at all, which is why this has to be allowed.
- **Alphabets are compared as sets.** Order and repeats in the slice mean nothing. `Intersect` and `Union` return an error when the alphabets differ. `Equivalent` and `Witness` have no error return, so they panic instead. The equivalence formula is built from `Complement`, which is relative to each machine's own alphabet, so it's only correct when both machines share one.
- **Every construction preserves its input's alphabet.** That includes a symbol whose transitions all lead to a dead state. `Compile` has no input machine to take an alphabet from, so the caller supplies `sigma`, and `Compile` builds over exactly that set, unused symbols included.
- **`Complement` is relative to the machine's own $\Sigma^*$.** A word containing a symbol outside $\Sigma$ is in neither L(a) nor its complement.

### Words and membership

- **A `Word` may contain any byte, and every membership test answers for every word.** A word containing a symbol outside the machine's alphabet isn't in $\Sigma^*$, so it isn't in the language: `Run` and `Contains` return false. They never panic. `engine.Regexp` has no alphabet, and the same applies to it: a word containing a symbol the pattern never mentions is rejected.
- **Acceptance is whole-word.** `Contains(w)` asks whether the language contains all of `w`, not whether `w` contains a match.
- **`Run` and `Contains` agree on every word**, and no call depends on any earlier one.
- **A nil `Word` and a zero-length `Word` are the same empty word.**

### States and validation

- **State labels are arbitrary `int`s.** They can be negative or non-contiguous, and the start state doesn't have to come first in the list.
- **A label listed twice in `states` is an error.** Q is where state labels are defined, so `{0, 1, 0}` is ambiguous. Repeats anywhere else only refer to something twice and are accepted with set semantics: in the alphabet, the accepting states, or an NFA's target lists.
- **The constructors reject malformed tuples**, and each constructor's doc lists what it rejects. The error text is up to you. When a constructor returns an error, nothing is promised about the machine it returns alongside it.
- **A DFA's transition function must be total**, and every DFA a construction returns is total too. An NFA's may be sparse: a missing entry or an empty target list means "no transition." ε-cycles, self-loops, and an empty accepting set are all legal.

### Observers

- **`NumStates()` is |Q|** as given to the constructor, or as built by a construction, counting unreachable and dead states. `DFA` and `NFA` both have it.
- **`DFA.Alphabet()` returns $\Sigma$ as a set**: each symbol exactly once, in any order, in a fresh slice that belongs to the caller. Changing that slice must not change the machine. `NFA` has no alphabet observer.

### Regular expressions

- **`Seq()` with no arguments denotes {ε}, and `Or()` with no arguments denotes ∅.** The empty concatenation and the empty union are the identities of their operations. With operands, `Seq` and `Or` use all of them, in order.
- **`String()` renders course regex syntax**, which Project 2's parser will read:

  | AST | Renders as |
  |---|---|
  | `Lit(c)` | the letter or digit itself |
  | `Eps()`, `Seq()` | `ε` |
  | `Emptyset()`, `Or()` | `∅` |
  | `Seq(a, b, ...)` | the operands' renderings concatenated, in order |
  | `Or(a, b, ...)` | the alternatives separated by `\|` |
  | `Star(a)` | postfix `*` on the whole operand |

  Operator precedence, from high to low, is: star, concatenation, then union. Extra parentheses are always allowed, so parenthesize defensively: `Seq(Or(Lit('a'), Lit('b')), Lit('c'))` is `(a|b)c`, not `a|bc`. There are no escapes and no whitespace. Write a nested star as `(a*)*`, never `a**`, and the empty word as `ε`, never `()` or an empty string.
- **`String()` is graded.** It's the only observable output of phase 1.5. The grading tests converts your String() output into a Go regular expression and checks its behavior against Go's `regexp` engine. 

### Arguments and results

- **No contract function modifies its arguments.** After any call, each machine argument behaves exactly as before: the same `Run`/`Contains` answers, the same `NumStates()`, and for a DFA the same `Alphabet()`. Some algorithms are described as editing a machine in place (state elimination rips out states, minimization merges them). Do that to a copy. The reason is the test oracles: in `Equivalent(Minimize(d), d)`, Go evaluates `Minimize(d)` before it reads `d`, so a minimizer that rewrites `d` gets compared with its own output and passes even when it's wrong.
- **Constructions are deterministic in size.** Repeating a construction on the same arguments gives a machine with the same `NumStates()`. For `FromDFA`, it gives an equally correct expression. A count that varies from call to call is a bug, even when every result decides the right language.
- **`NFAToDFA` builds only reachable subsets:** at most one state per subset of the NFA's states reachable from the start subset, plus possibly the empty subset. Fewer is fine; this is a bound, not a minimality requirement.
- **`Node`s can be shared**, more than once in one tree or across several trees. `String()`, `Compile`, and `engine.FromAST` must give a shared node the same meaning at every occurrence and return the same result on every call.
- **Results may share internal structure with their inputs**, since nothing ever writes to either.

### Preconditions and ownership

- **Arguments come from the contract.** A machine argument must come from `NewDFA`, `NewNFA`, or a construction. A `Node` must come from the six AST constructors, with no nil children. Passing nil, or a machine assembled any other way, is the caller's mistake: behavior is unspecified, a panic is fine, and your code doesn't need to check. The tests never do it.
- **Nil slices, maps, and `Word`s are empty values.** A nil `Word` is the empty word, a nil `sigma` is the empty alphabet, and a nil `delta` or `epsilon` means no transitions (legal for an NFA, and legal for a DFA only when there is no state-symbol pair to cover).
- **Constructors don't have to copy their arguments.** Callers, including the tests, don't modify a slice or map after passing it to a constructor. The one ownership promise runs the other way: the slice `Alphabet()` returns belongs to the caller.

### Unspecified

Anything not stated here or in a function's doc comment is unspecified. That includes your representations, error message text, the state labels a construction picks, the exact shape of the expression `FromDFA` returns (any correct one will do), which shortest word `Witness` returns when there are several, and performance beyond finishing promptly.

## Design considerations

Warnings and hints that span more than one function:

- **Build every machine through `NewDFA` or `NewNFA`**, including every machine a construction returns, rather than assembling a struct by hand. The constructors' checks then catch a construction bug at the line that caused it. `Complement`, `NFAToDFA`, `Minimize`, and `DFAToNFA` have no error return, so if the constructor returns an error inside one of them, `panic` with it. That error is your bug, and discarding it hands the caller a nil machine that crashes somewhere unrelated.
- **Plan at phase 1.1 for reading a DFA from another package.** At phase 1.7, `regex.FromDFA` has to read a DFA's states, transitions, start state, and accepting states. Go only lets `regex` see what `automata` exports, and the contract exports none of those, so you'll need to expose them yourself, in whatever shape suits your representation. Use exported *methods* (`d.Start()`, `d.Step(q, a)`, and so on) rather than exported struct fields. Exported fields would let any package build a `DFA` that never passed through `NewDFA`, or change one afterwards, and the rest of the contract relies on every `*DFA` being valid. That's why `Complement` has no error return.
- **Watch for the zero value on map lookups.** Indexing a transition map with a key it doesn't have returns the zero `State`, which is often a real state (usually the start). A DFA `Run` that just indexes the map doesn't reject a foreign symbol.
- **Go map iteration order is randomized.** If you identify a set of states by a key built while ranging over a map, the same set can get two different keys. Sort before building a key. This is the usual way subset construction and partition refinement break the determinism rule.
- **`FromDFA` should leave out symbols that only lead to a dead state.** "No `c`" over {a, b, c} is `(a|b)*`. The expression then no longer mentions `c`, so a machine compiled over only the symbols it mentions has a smaller alphabet than `d`, and `Equivalent` panics on machines whose alphabets differ. When you check `FromDFA`'s output, compile it over `d`'s full alphabet, which brings the dropped symbol back as a dead state.
- **Don't write a backtracking matcher.** Trying alternative ways to split the input is exponential on patterns like `(a|a)*b` against a long run of `a`s, which the grading suite includes. 
- **Scale.** The grading suites include some large inputs, and every check runs under a time limit (the handout lists both). A correct implementation finishes each one in milliseconds. Recursion as deep as the input (one level per symbol of a long word, or per level of a deep AST) is fine, because Go stacks grow as needed. Polynomial algorithms are fine too. You should only need to watch out for infinite recursion and exponential runtimes.

## Testing

You're responsible for your own testing. 

**Prefer deciding to sampling.** Once `Equivalent` works, most correctness claims in this project are decidable, so decide them. `Witness` hands you the shortest counterexample when they fail. Fall back on generated words only where no exact check exists. You can build many languages as a closure expression over named machines. A test of your own, in `tests/`, might look like this:

```go
func TestP1_2_EvenZerosNotEndingIn01(t *testing.T) {
    lang := oracle.Must(automata.Intersect(
        corpus.EvenZeros(), automata.Complement(corpus.EndsWith01())))
    oracle.Accepts(t, lang, "0011") // two 0s, doesn't end in 01
    oracle.Rejects(t, lang, "0101") // two 0s, but ends in 01
}
```

**Exact checks by phase.** Each construction you build comes with a theorem about the language of what it produces, and each of those theorems is a claim you can decide. Turning them into tests is part of the work, and a good check on whether you understand what each construction is for. Theorem and lemma numbers are Sipser's.

| Phase | What the theory guarantees |
|---|---|
| 1.4 | The subset construction (Theorem 1.39) gives a DFA that recognizes the same language as the NFA it starts from. Every DFA can also be viewed as an NFA without changing its language. |
| 1.6 | Thompson's construction (Lemma 1.55) gives an NFA that recognizes exactly $L(R)$, so it recognizes the same language as any DFA you know is correct for $L(R)$. |
| 1.7 | State elimination (Lemma 1.60) gives a regular expression $R$ with $L(R) = L(D)$. Together with phase 1.6, that takes you from a DFA to a regex and back to a machine for the same language. |
| 1.8 | Minimization preserves the language. The minimal DFA for a language over $\Sigma$ is unique up to renaming states (Myhill-Nerode, Problem 1.52), so its state count depends only on the language. The corpus lists that count for each of its machines. A DFA that's already minimal can't get any smaller. |
| 1.9 | Matching answers membership in $L(R)$, and every correct regex implementation gives the same answers. Go's `regexp` package is one. The engine has no exact check, so this is where generated words come in. |

**Provided packages** (`go doc -all ./oracle` and `go doc -all ./corpus` for details):

| Name | Use |
|---|---|
| `oracle.Accepts`, `Rejects`, `Decides` | assert a decision on one word; they take any `automata.Language` and a plain Go `string` |
| `oracle.Must` | drop an `error` return so closure operations compose in one expression |
| `oracle.AssertEquivalent` | assert two DFAs have the same language, reporting the shortest distinguishing word |
| `oracle.CheckAgainstStdlib` | compare any `Language` with Go's `regexp` on an AST, anchored and translated for you |
| `oracle.GoPattern` | translate course syntax (`ε`, `∅`) to Go syntax; anchor with `\A(?:` ... `)\z` before whole-word matching |
| `oracle.NewGen` | seeded random words (`Word`, `Words`) and ASTs (`AST`) |
| `corpus.All()` | every named machine, with its name and its minimal state count (`MinStates`) |
| `corpus.Sigma` | the corpus machines' alphabet |

## Later projects

Project 2, and possibly later projects, add new files and packages to this repository. They never change a file that has already shipped, so nothing in this README changes.
