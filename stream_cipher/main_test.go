// Author: Arkan M. Gerges<arkan.m.gerges@gmail.com>

package stream_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptAndDecryptAreEqual(t *testing.T) {
	plaintext := "This is a secret and you should not be able to read it 0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	key := NewLcg(10)
	ciphertext := Process(key, []byte(plaintext))
	key = NewLcg(10)
	assert.Equal(t, plaintext, string(Process(key, ciphertext)))
}

func TestEncryptAndDecryptAreNotEqualWhenKeysAreDifferent(t *testing.T) {
	plaintext := "This is a secret and you should not be able to read it 0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	key := NewLcg(10)
	ciphertext := Process(key, []byte(plaintext))
	key = NewLcg(11)
	assert.NotEqual(t, plaintext, string(Process(key, ciphertext)))
}
