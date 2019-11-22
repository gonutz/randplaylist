package randplaylist

import (
	"testing"

	"github.com/gonutz/check"
)

func TestIsPrime(t *testing.T) {
	primes := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67,
		71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139,
		149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199,
	}
	contains := func(list []int, n int) bool {
		for _, m := range list {
			if n == m {
				return true
			}
		}
		return false
	}
	for i := -1; i <= 200; i++ {
		check.Eq(t, isPrime(i), contains(primes, i), i)
	}
}

func TestNextPrime(t *testing.T) {
	check.Eq(t, nextPrime(-1), 2)
	check.Eq(t, nextPrime(0), 2)
	check.Eq(t, nextPrime(1), 2)
	check.Eq(t, nextPrime(2), 3)
	check.Eq(t, nextPrime(3), 5)
	check.Eq(t, nextPrime(4), 5)
	check.Eq(t, nextPrime(5), 7)
}

func TestTrivialSizes(t *testing.T) {
	r := New(-1)
	for i := 0; i < 100; i++ {
		check.Eq(t, r.Next(), -1)
	}

	r = New(0)
	for i := 0; i < 100; i++ {
		check.Eq(t, r.Next(), -1)
	}

	r = New(1)
	for i := 0; i < 100; i++ {
		check.Eq(t, r.Next(), 0)
	}
}

func TestSizeTwoTogglesBetween0and1(t *testing.T) {
	for i := 0; i < 100; i++ {
		r := New(2)
		// Make sure the next number generated will be 0 so we can test it. If
		// r.Next() returns 1 first, we expected 0 next which is what we want.
		// If it returns 0 first, we call Next once more (should return 1) and
		// expect to see 0 afterwards.
		if r.Next() == 0 {
			r.Next()
		}
		for j := 0; j < 100; j++ {
			check.Eq(t, r.Next(), 0)
			check.Eq(t, r.Next(), 1)
		}
	}
}

func TestRandSetGoesThroughWholeSequenceBeforeRestarting(t *testing.T) {
	for i := 0; i < 100; i++ {
		r := New(10)
		var seen [10]int
		for n := 1; n < 10; n++ {
			for range seen {
				seen[r.Next()]++
			}
			// Generating the next 10 items should have each item incremented
			// once.
			check.Eq(t, seen, [10]int{n, n, n, n, n, n, n, n, n, n})
		}
	}
}

func TestLastNumberDoesNotRepeatAfterSequence(t *testing.T) {
	// Once every index has been generated we want to make sure that the next
	// index directly after that is not the same as the end of the last
	// sequence. Otherwise our playlist might play the same song twice in a row
	// after all tracks have been played.
	r := New(3)
	last := r.Next()
	for i := 0; i < 1000; i++ {
		next := r.Next()
		if last == next {
			t.Fail()
		}
		last = next
	}
}
