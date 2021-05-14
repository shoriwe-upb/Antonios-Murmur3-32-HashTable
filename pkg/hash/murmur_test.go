package hash

import (
	"log"
	"testing"
)

func init() {
	log.Println("Testing Murmur32 Hash")
}

const samplesSeek uint32 = 100

var samples = map[string]uint32{

}

func TestMurmur3_32Bits(t *testing.T) {
	for sample, expectedResult := range samples {
		computedHash := Murmur3_32Bits(sample, samplesSeek)
		if expectedResult != computedHash {
			t.Errorf("Received invalid Murmur32 Hash for %s, expecting %d, but received %d", sample, expectedResult, computedHash)
			return
		}
	}
	t.Logf("Murmur32 Hash Success")
}
