package siphash

import (
	"encoding/binary"

	dchest_siphash "github.com/dchest/siphash"
)

func Fuzz(input []byte) {
	var k0, k1 uint64
	if len(input) > 16 {
		k0 = binary.LittleEndian.Uint64(input[0:8])
		k1 = binary.LittleEndian.Uint64(input[8:16])
	}
	dchest_siphash.Hash(k0, k1, input)
	dchest_siphash.Hash128(k0, k1, input)
}
