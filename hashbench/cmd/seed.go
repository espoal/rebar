package main

import (
	"github.com/espoal/rebar/hashbench"
	"log"
)

func main() {

	size := 4

	key := []byte("hello world")
	buff := make([]byte, size*16)
	seeds := hashbench.SeedFactory(0x1234567890abcdef, size)

	hashbench.Xxh3WithSeed(key, buff, size, seeds)

	log.Println("buff", buff)

}
