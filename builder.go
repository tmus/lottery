package lottery

import (
	"math/rand"
)

type builder struct {
	odd int
	in  int

	success func()
	failure func()

	rand *rand.Rand
}

// Success sets a function on the lottery builder that will run if the lottery
// odds are met.
func (b *builder) Success(fn func()) *builder {
	b.success = fn
	return b
}

// Failure sets a function on the lottery builder that will run if the lottery
// odds are _not_ met.
func (b *builder) Failure(fn func()) *builder {
	b.failure = fn
	return b
}

// Choose generates a random number and determines if the result satisfies odds
// for the lottery. If so, the Success function is ran. If not, the Failure
// function is ran.
func (b *builder) Choose() {
	b.ChooseMany(1)
}

func (b *builder) ChooseMany(times int) {
	for i := 0; i < times; i++ {
		b.choose()
	}
}

func (b *builder) choose() {
	if b.wins() {
		b.success()
	} else {
		b.failure()
	}
}

func (b *builder) wins() bool {
	return b.rand.Intn(b.in) < b.odd
}

type resultBuilder[T any] struct {
	odd int
	in  int

	success func() T
	failure func() T

	rand *rand.Rand
}

func (b *resultBuilder[T]) Odds(odd int, in int) *resultBuilder[T] {
	b.odd = odd
	b.in = in
	return b
}

func (b *resultBuilder[T]) Success(fn func() T) *resultBuilder[T] {
	b.success = fn
	return b
}

func (b *resultBuilder[T]) Choose() T {
	if b.wins() {
		return b.success()
	} else {
		return b.failure()
	}
}

func (b *resultBuilder[T]) wins() bool {
	return b.rand.Intn(b.in) < b.odd
}
