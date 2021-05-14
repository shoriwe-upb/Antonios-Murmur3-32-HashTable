package hash

import (
	"encoding/binary"
)

/*
	Murmur32 Hash implementation based on the C one found at
	https://en.wikipedia.org/wiki/MurmurHash
*/

func Murmur32Scramble(k uint32) uint32 {
	k *= 0xcc9e2d51
	k = (k << 15) | (k >> 17)
	k *= 0x1b873593
	return k
}

// Murmur3_32Bits -> Murmur hash
func Murmur3_32Bits(value string, seek uint32) uint32 {
	valueLength := uint32(len(value))

	h := seek
	var k uint32

	index := uint32(0)
	for i := valueLength >> 2; i != 0; i-- {
		a := value[index]
		index++
		b := value[index]
		index++
		c := value[index]
		index++
		d := value[index]
		index++

		k = binary.LittleEndian.Uint32([]uint8{a, b, c, d})

		h ^= Murmur32Scramble(k)
		h = (h << 13) | (h >> 19)
		h = h*5 + 0xe6546b64
	}
	k = 0
	for currentIndex := valueLength - 1; currentIndex >= index; currentIndex-- {
		k <<= 8
		k |= uint32(value[currentIndex])
	}

	h ^= Murmur32Scramble(k)

	h ^= valueLength
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}
