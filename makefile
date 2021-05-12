.PHONY: bench-all
bench-all:
	go test -bench=. ./src

.PHONY: static-vs-oo
static-vs-oo:
	go test -bench=. ./src/static_vs_oo_test.go
