// Author: Arkan M. Gerges<arkan.m.gerges@gmail.com>
// The Linear Congruential Generator (LCG) is a simple pseudorandom number generator.
// See https://en.wikipedia.org/wiki/Linear_congruential_generator

package stream_cipher

import (
	"math"
)

type Lcg struct {
	next    uint
	modulus uint
}

func NewLcg(seed uint) *Lcg {
	return &Lcg{seed, uint(math.Pow(2, 31))}
}

func (lcg *Lcg) NextNumber() uint {
	// The following numbers were taken from the Wikipedia article on LCGs (see the link above), and uses the ANSI C parameters.
	lcg.next = (lcg.next*1103515245 + 12345) % lcg.modulus
	return lcg.next
}

func (lcg *Lcg) NextNumberMod256() uint {
	return lcg.NextNumber() % 256
}

func Process(key *Lcg, message []byte) []byte {
	result := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		result[i] = message[i] ^ byte(key.NextNumberMod256())
	}

	return result
}
