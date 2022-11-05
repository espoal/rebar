package main

import (
	"github.com/espoal/rebar/hashbench"
	"github.com/zeebo/blake3"
	"log"
)

func main() {

	size := 4

	key := []byte("hello world")
	buff := make([]byte, size*16)
	h := blake3.New()

	hashbench.B3(key, buff, h)

	log.Println("buff", buff)

}
