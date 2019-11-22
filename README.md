Inspired by a chapter in Mike McShaffry's book Game Coding Complete, this library provides a function to generate a random sequence of indices [0..N] where each index is generated exactly once before the playlist repeats.

Example:

```
package main

import (
	"fmt"
	"github.com/gonutz/randplaylist"
)

func main() {
	r := randplaylist.New(5)
	for i := 0; i < 5; i++ {
		fmt.Print(r.Next(), " ")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print(r.Next(), " ")
	}
}
```
