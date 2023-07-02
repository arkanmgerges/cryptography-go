// Author: Arkan M. Gerges<arkan.m.gerges@gmail.com>

package substitution_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptAndDecryptAreEqual(t *testing.T) {
	plaintext := "This is a secret and you should not be able to read it 0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	key, err := GenerateKey()
	if err != nil {
		t.Error(err)
	}
	ciphertext := Encrypt(key, plaintext)
	assert.Equal(t, plaintext, Decrypt(key, ciphertext))
}

func TestEncryptAndDecryptAreNotEqual(t *testing.T) {
	plaintext := "This is a secret and you should not be able to read it 0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	key, err := GenerateKey()
	key2, err2 := GenerateKey()
	if err != nil || err2 != nil {
		t.Error(err)
	}

	ciphertext := Encrypt(key, plaintext)
	assert.NotEqual(t, plaintext, Decrypt(key2, ciphertext))
}
