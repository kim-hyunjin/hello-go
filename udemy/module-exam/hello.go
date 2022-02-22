package hello

import (
	// "rsc.io/quote"
	"rsc.io/quote/v3" // Each different major version (v1, v2, and so on) of a Go module uses a different module path
)

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
    return quote.Concurrency()
}