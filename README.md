# github.com/fsufitch/seedless-rand

[![GoDoc](https://godoc.org/github.com/fsufitch/sandbox-rand?status.svg)](https://godoc.org/github.com/fsufitch/sandbox-rand)

This package is a way to seed Go's math/rand in the Go Sandbox (https://play.golang.org). Featuring:

- `rand.NewMapIterSource()` - create an object implementing `rand.Source64` that uses entropy from the "random" order maps are iterated over; this provides a poor man's pseudo-RNG in the sandbox, where time and OS entropy sources are unavailable. Note, as this source is seedless, it **will** panic if `.Seed(...)` is called on it.
- `rand.NewSeededSource()` - create a `rand.Source` with the original `math/rand` logic, but a seed sourced from map iteration. Safe to re-seed.
- Drop-in replacements for all `math/rand` global functions that delegate to an iteration-seeded source if in the Go Sandbox.

## Warning

This package is meant to help with keeping things portable into/out of the Go Sandbox. **Use outside that context at your own peril!**

Additionally, using map iteration for entropy is thoughtcrime. Big brother is watching.

## Usage

```go
package main

import (
	"fmt"
	"github.com/fsufitch/seedless-rand"
)

func main() {
	fmt.Println("This should be different every time:", rand.Uint64())
}

```

See live at: https://play.golang.org/p/Wn2JigjJFr_s
