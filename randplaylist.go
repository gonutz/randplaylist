package randplaylist

import "math/rand"

// Playlist generates sequences of indices that do not repeat until every index
// was returned once by the Next function.
type Playlist interface {
	Next() int
}

// New returns a new Playlist with the math/rand.Intn function for generating
// the random numbers used in the algorithm. See NewWithRand for details.
func New(size int) Playlist {
	return NewWithRand(size, rand.Intn)
}

// NewWithRand returns a Playlist that generates a sequence of pseudo-random
// numbers in the range [0..size-1] in which each number is only returned once
// by the Playlist.Next function until all numbers in that range have been
// returned. After one such cycle the randN function is used to re-seed the
// algorithm to return a different sequence each cycle.
//
// The randN function must return a random value in [0..n) (excluding n).
//
// If size is <= 0 the Next function will always return -1.
func NewWithRand(size int, randN func(n int) int) Playlist {
	if size <= 0 {
		return constant(-1)
	}
	if size == 1 {
		return constant(0)
	}
	if size == 2 {
		r := toggle(randN(2))
		return &r
	}
	r := &playlist{
		size:  size,
		prime: nextPrime(size),
		randN: randN,
	}
	r.seed()
	return r
}

type constant int

func (c constant) Next() int {
	return int(c)
}

type toggle int

func (t *toggle) Next() int {
	*t = 1 - *t
	return int(*t)
}

func nextPrime(n int) int {
	n++
	for !isPrime(n) {
		n++
	}
	return n
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

type playlist struct {
	size      int
	prime     int
	skip      int
	cur       int
	remaining int
	randN     func(n int) int
}

func (r *playlist) seed() {
	rgn := func() int { return 1 + r.randN(r.size) }
	a, b, c := rgn(), rgn(), rgn()
	r.skip = a*r.size*r.size + b*r.size + c
	// The random skip value might end up generating the same value all over. In
	// that case we just increment it, this will break the circle.
	if r.skip%r.prime == 0 {
		r.skip++
	}
	r.remaining = r.size + 1 // +1 because Next will decrement it
	r.Next()
}

func (r *playlist) Next() int {
	if r.remaining <= 0 {
		r.seed()
	}
	r.remaining--

	r.cur = (r.cur + r.skip) % r.prime
	for r.cur >= r.size {
		r.cur = (r.cur + r.skip) % r.prime
	}
	return r.cur
}
