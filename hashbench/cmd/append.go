package main

import (
	"github.com/espoal/rebar/hashbench"
	"github.com/zeebo/xxh3"
	"log"
)

func main() {

	size := 2

	key := []byte("hello world")
	buff := make([]byte, size*16)
	seeds := hashbench.RandBytes(size, 16)

	h := xxh3.New()

	hashbench.Xxh3Append(key, buff, size, seeds, h)

	log.Println("buff", buff)

}
