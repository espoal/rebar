package hashbench

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/zeebo/xxh3"
)

func SeedFactory(seed uint64, size int) []uint64 {

	buff := make([]uint64, 1, size)
	buff[0] = seed

	for i := 1; i < size; i++ {
		next := xxh3.Hash(binary.BigEndian.AppendUint64(nil, buff[i-1]))
		buff = append(buff, next)
	}

	return buff
}

func RandBytes(count, length int) [][]byte {
	b := make([][]byte, count)
	for i := range b {
		b[i] = make([]byte, length)
		_, err := rand.Read(b[i])
		if err != nil {
			return nil
		}
	}
	return b
}

func Xxh3WithSeed(key, buff []byte, size int, seed []uint64) []byte {

	//buff := make([]byte, 0, size*16)

	for i := 0; i < size; i++ {
		hash := xxh3.Hash128Seed(key, seed[i]).Bytes()
		copy(buff[i*HASH_LEN:], hash[:])
	}

	return buff
}

func Xxh3Append(key, buff []byte, size int, seeds [][]byte, h *xxh3.Hasher) []byte {
	
	h.Reset()
	_, err := h.Write(key)
	if err != nil {
		panic(err)
	}

	for i := 0; i < size; i++ {
		hash := h.Sum128().Bytes()
		//buff = append(buff, hash[:]...)
		copy(buff[i*16:], hash[:])
		_, err = h.Write(seeds[i])
		if err != nil {
			panic(err)
		}

	}

	return buff
}
