// Author: Arkan M. Gerges<arkan.m.gerges@gmail.com>

package caesar_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptAndDecryptWithKeyOffset3(t *testing.T) {
	plaintext := "This is a secret and you should not be able to read it 0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	assert.Equal(t, plaintext, Decrypt(3, Encrypt(3, plaintext)))
}

func TestEncryptAndDecryptWithKeyOffset30(t *testing.T) {
	plaintext := "This is a secret and you should not be able to read it 0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	assert.Equal(t, plaintext, Decrypt(30, Encrypt(30, plaintext)))
}

func TestEncryptAndDecryptWithKeyOffset100(t *testing.T) {
	plaintext := "This is a secret and you should not be able to read it 0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	assert.Equal(t, plaintext, Decrypt(100, Encrypt(100, plaintext)))
}

func TestEncryptAndDecryptWithDifferentKeys(t *testing.T) {
	plaintext := "This is a secret and you should not be able to read it 0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	assert.NotEqual(t, plaintext, Decrypt(30, Encrypt(100, plaintext)))
}
