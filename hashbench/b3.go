package hashbench

import "github.com/zeebo/blake3"

func B3(key, buff []byte, h *blake3.Hasher) []byte {

	h.Reset()
	_, err := h.Write(key)
	if err != nil {
		panic(err)
	}
	d := h.Digest()
	_, err = d.Read(buff)
	if err != nil {
		panic(err)
	}

	return nil
}
