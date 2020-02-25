package rand

import "math/rand"

var bits = map[bool]struct{}{false: struct{}{}, true: struct{}{}}

// iterSource is a rand.Source64 that produces values using "random" map iteration order
type iterSource struct{}

// Uint64 generates a "random" number [0..2**64) using "random" map iteration order
func (s iterSource) Uint64() uint64 {
	var result uint64
	for i := 0; i < 64; i++ {
		for b := range bits {
			result = (result << 1)
			if b {
				result |= 1
			}
			break
		}

	}
	return result
}

// Int63 generates a "random" number [0..2**63) using "random" map iteration order
func (s iterSource) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

// Seed panics; it's impossible to seed random map iteration
func (s iterSource) Seed(int64) {
	panic("seeding a seedless source is not supported")
}

// NewMapIterSource returns a new rand.Source64 that usese map iteration for entropy.
// This source is seedless, stateless and will panic if Seed is called
func NewMapIterSource() rand.Source64 {
	return iterSource{}
}

// NewSeededSource creates a new rand.Source, pre-seeded using map iteration
func NewSeededSource() rand.Source {
	return NewSource(iterSource{}.Int63())
}
