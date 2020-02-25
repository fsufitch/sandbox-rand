package rand

import "math/rand"

// selectSource is a rand.Source64 that produces values using "random" select mechanism
type selectSource struct{}

// Uint64 generates a "random" number [0..2**64) using "random" select mechanism
// More here: https://golang.org/ref/spec#Select_statements
func (s selectSource) Uint64() (x uint64) {
	ch := make(chan struct{})
	close(ch)
	for i := 0; i < 64; i++ {
		select {
		case <-ch:
			x = x<<1 | 1
		case <-ch:
			x = x<<1 | 0
		}
	}
	return
}

// Int63 generates a "random" number [0..2**63)
func (s selectSource) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

// Seed panics; it's impossible to seed a seedless source
func (s selectSource) Seed(int64) {
	panic("seeding a seedless source is not supported")
}

// NewSeedlessSource returns a new rand.Source64 that uses selects for entropy.
// This source is seedless, stateless and will panic if Seed is called
func NewSeedlessSource() rand.Source64 {
	return selectSource{}
}

// NewSeededSource creates a new rand.Source, pre-seeded using NewSeedlessSource()
func NewSeededSource() rand.Source {
	return NewSource(selectSource{}.Int63())
}
