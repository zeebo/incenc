// +build ignore

package main

import (
	"math/rand"

	"github.com/dvyukov/go-fuzz/gen"
	"github.com/zeebo/incenc"
)

func main() {
	var buf []byte

	for {
		var w incenc.Writer
		buf = buf[:0]
		values := rand.Intn(10)
		for i := 0; i < values; i++ {
			buf, _ = w.AppendBytes(buf, randomValue())
		}
		gen.Emit(buf, nil, true)
	}
}

func randomValue() []byte {
	out := make([]byte, rand.Intn(10))
	for i := range out {
		out[i] = byte(rand.Intn(10))
	}
	return out
}
