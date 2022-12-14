package lottery

import (
	"math/rand"
	"time"
)

// Odds creates a new instance of a lottery with the provided odds. Chain calls
// to Success and Failure in order to pass any behaviour that should occur as
// a result of the lottery. Note both calls are optional: if a function is
// not provided for Success or Failure, a no-op is performed.
func Odds(odd int, in int) *builder {
	s := rand.NewSource(time.Now().UnixMicro())

	return &builder{
		odd: odd,
		in:  in,

		success: func() {},
		failure: func() {},

		rand: rand.New(s),
	}
}
