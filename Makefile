# Project 1 — regular-languages toolkit.
#
# The targets you will actually use are `test` and, at phase 1.9, `bench`.

.PHONY: all test smoke bench vet fmt clean

all: vet test

# All tests. Expect "unimplemented: phase 1.x ..."
# failures for phases you have not built yet — that is the expected state, not
# a broken checkout.
test:
	go test ./...

# The one thing that should pass on a fresh clone, before you write anything.
smoke:
	go test ./tests/ -run 'TestP1_0_Types$$' -v

# Phase 1.9. The handout asks you to report these numbers.
bench:
	go test ./... -run '^$$' -bench . -benchmem

vet:
	go vet ./...

fmt:
	gofmt -l -w .

clean:
	go clean -testcache
