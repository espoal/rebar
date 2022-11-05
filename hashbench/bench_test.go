package hashbench

import (
	"github.com/zeebo/blake3"
	"github.com/zeebo/xxh3"
	"testing"
)

func BenchmarkXXH3(b *testing.B) {

	strings := RandBytes(b.N, STRINGS_LEN)

	seeds := SeedFactory(SEED, SIZE)
	buff := make([]byte, HASH_SIZE)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Xxh3WithSeed(strings[n], buff, SIZE, seeds)
	}

}

func BenchmarkXXH3Append(b *testing.B) {

	strings := RandBytes(b.N, STRINGS_LEN)

	seeds := RandBytes(SIZE, HASH_LEN)
	buff := make([]byte, HASH_SIZE)
	h := xxh3.New()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Xxh3Append(strings[n], buff, SIZE, seeds, h)
	}

}

func BenchmarkB3(b *testing.B) {

	strings := RandBytes(b.N, STRINGS_LEN)

	buff := make([]byte, HASH_LEN)
	h := blake3.New()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		B3(strings[n], buff, h)
	}

}
