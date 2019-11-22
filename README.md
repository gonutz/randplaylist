Inspired by a chapter in Mike McShaffry's book Game Coding Complete, this library provides a function to generate a random sequence of indices `[0..N]` where each index is generated exactly once before the playlist repeats.

Example:

```Go
package main

import (
	"fmt"
	"github.com/gonutz/randplaylist"
)

func main() {
	const (
		trackCountInAlbum = 5
		albumRepeatCount  = 3
	)

	r := randplaylist.New(trackCountInAlbum)
	for i := 0; i < albumRepeatCount; i++ {
		for j := 0; j < trackCountInAlbum; j++ {
			fmt.Print(r.Next(), " ")
		}
		fmt.Println()
	}
}

// Example output:
// 1 4 2 0 3
// 1 0 4 3 2
// 0 4 1 2 3
```

The algorithm works as follows:

Given the number of tracks in the playlist `N`, find the next prime number greater than `N` (not `N` itself if it is prime). Call this `prime`.

Now we need three random numbers `a, b, c` to generate a `skip` value

    skip = a*N*N + b*N + c

To now find the next index in the sequence we can use this formula:

    index = (index + skip) % prime

This is the basic principle. We have to make sure to generate a new index if it lies outside the track range because the term `... % prime` can result in values `[0..prime)` so we might get a value that is `>= N`. Thus we need to repeat the index calculation until we get a valid value.
