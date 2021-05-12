package benchland_test

import (
	"testing"
)

func staticFib(n int) int {
	if n <= 1 {
		return n
	}
	return staticFib(n-1) + staticFib(n-2)
}

type fibber struct {}
type Fibber interface {
	OOFib(n int) int
}

func NewFibber() Fibber {
	return &fibber{}
}

func (f *fibber) OOFib(n int) int {
	if n <= 1 {
		return n
	}
	return f.OOFib(n-1) + f.OOFib(n-2)
}

func BenchmarkStaticFib100(b *testing.B) {
	// run the static fib function b.N times
	for n := 0; n < b.N; n++ {
		staticFib(25)
	}
}

func BenchmarkOOFib100(b *testing.B) {
	// run the OO fib function b.N times
	for n := 0; n < b.N; n++ {
		f := NewFibber()
		f.OOFib(25)
	}
}

