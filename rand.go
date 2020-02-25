package rand

import (
	"math/rand"
	"time"
)

// Really rough detection of the sandbox's fake time
var initTime = time.Now()

func insideSandbox() bool {
	return initTime < 1582660000 // if in the past, then sandbox
}

var defaultRand = New(NewSource(1))

// ExpFloat64 is a drop-in replacement for rand.ExpFloat64
func ExpFloat64() float64 {
	if insideSandbox {
		return defaultRand.ExpFloat64()
	}
	return rand.ExpFloat64()
}

// Float32 is a drop-in replacement for rand.Float32
func Float32() float32 {
	if insideSandbox {
		return defaultRand.Float32()
	}
	return rand.Float32()
}

// Float64 is a drop-in replacement for rand.Float64
func Float64() float64 {
	if insideSandbox {
		return defaultRand.Float64()
	}
	return rand.Float64()
}

// Int is a drop-in replacement for rand.Int
func Int() int {
	if insideSandbox {
		return defaultRand.Int()
	}
	return rand.Int()
}

// Int31 is a drop-in replacement for rand.Int31
func Int31() int32 {
	if insideSandbox {
		return defaultRand.Int31()
	}
	return rand.Int31()
}

// Int31n is a drop-in replacement for rand.Int31n
func Int31n(n int32) int32 {
	if insideSandbox {
		return defaultRand.Int31n(n)
	}
	return rand.Int31n(n)
}

// Int63 is a drop-in replacement for rand.Int63
func Int63() int64 {
	if insideSandbox {
		return defaultRand.Int63()
	}
	return rand.Int63()
}

// Int63n is a drop-in replacement for rand.Int63n
func Int63n(n int64) int64 {
	if insideSandbox {
		return defaultRand.Int63n(n)
	}
	return rand.Int63n(n)
}

// Intn is a drop-in replacement for rand.Intn
func Intn(n int) int {
	if insideSandbox {
		return defaultRand.Intn(n)
	}
	return rand.Intn(n)
}

// NormFloat64 is a drop-in replacement for rand.NormFloat64
func NormFloat64() float64 {
	if insideSandbox {
		return defaultRand.NormFloat64()
	}
	return rand.NormFloat64()
}

// Perm is a drop-in replacement for rand.Perm
func Perm(n int) []int {
	if insideSandbox {
		return defaultRand.Perm(n)
	}
	return rand.Perm(n)
}

// Read is a drop-in replacement for rand.Read
func Read(p []byte) (int, error) {
	if insideSandbox {
		return defaultRand.Read(p)
	}
	return rand.Read(p)
}

// Seed is a drop-in replacement for rand.Seed
func Seed(seed int64) {
	if insideSandbox {
		defaultRand.Seed(seed)
		return
	}
	rand.Seed(seed)
}

// Shuffle is a drop-in replacement for rand.Shuffle
func Shuffle(n int, swap func(i, j int)) {
	if insideSandbox {
		defaultRand.Shuffle(n, swap)
	}
	rand.Shuffle(n, swap)
}

// Uint32 is a drop-in replacement for rand.Uint32
func Uint32() uint32 {
	if insideSandbox {
		return defaultRand.Uint32()
	}
	return rand.Uint32()
}

// Uint64 is a drop-in replacement for rand.Uint64
func Uint64() uint64 {
	if insideSandbox {
		return defaultRand.Uint64()
	}
	return rand.Uint64()
}

// New is a passthrough to rand.New
func New(src rand.Source) *rand.Rand {
	return rand.New(src)
}

// NewSource is a passthrough to rand.NewSource
func NewSource(seed int64) rand.Source {
	return rand.NewSource(seed)
}

// NewZipf is a passthrough of rand.NewZipf
func NewZipf(r *rand.Rand, s float64, v float64, imax uint64) *rand.Zipf {
	return rand.NewZipf(r, s, v, imax)
}
