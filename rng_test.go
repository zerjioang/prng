package prng

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	generator = rngSource{}
)

func init() {
	generator.Seed(time.Now().Unix())
	rand.Seed(time.Now().Unix())
}

func TestRng(t *testing.T) {
	t.Run("uuid-0", func(t *testing.T) {
		fmt.Println(gen())
	})
	t.Run("uuid-1", func(t *testing.T) {
		fmt.Println(generator.UUID())
	})
	t.Run("uuid-2", func(t *testing.T) {
		var dst RandUUID
		generator.UUID2(&dst)
		fmt.Println(string(dst[:]))
		generator.UUID2(&dst)
		fmt.Println(string(dst[:]))
	})
	t.Run("uuid-3", func(t *testing.T) {
		fmt.Println(New())
	})
}

// gen is a simple 8 bytes hex encoded string generator
func gen() string {
	dst := make([]byte, 8)
	for i := 0; i < 8; i++ {
		dst[i] = byte(rand.Intn(256))
	}
	return hex.EncodeToString(dst[:])
}

func BenchmarkRng(b *testing.B) {
	b.Run("uuid-0", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = gen()
		}
	})
	b.Run("uuid-1", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = generator.UUID()
		}
	})
	b.Run("uuid-2", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		var dst RandUUID
		for i := 0; i < b.N; i++ {
			generator.UUID2(&dst)
		}
	})
	b.Run("uuid-3", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			New()
		}
	})
}

func ExampleNew() {
	New()
}
