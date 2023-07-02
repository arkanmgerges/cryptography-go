// Author: Arkan M. Gerges<arkan.m.gerges@gmail.com>

package substitution_cipher

import (
	"crypto/rand"
	"math/big"
)

func Encrypt(key map[string]string, plaintext string) string {
	result := ""
	for _, char := range plaintext {
		if key[string(char)] == "" {
			result += string(char)
		} else {
			result += key[string(char)]
		}
	}
	return result
}

func Decrypt(key map[string]string, ciphertext string) string {
	result := ""
	decryptionKey := GenerateDecryptionKey(key)
	for _, char := range ciphertext {
		if decryptionKey[string(char)] == "" {
			result += string(char)
		} else {
			result += decryptionKey[string(char)]
		}
	}
	return result
}

func GenerateKey() (map[string]string, error) {
	result := make(map[string]string)
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	tmpLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"

	for _, char := range letters {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(tmpLetters))))
		if err != nil {
			return nil, err
		}
		result[string(char)] = string(tmpLetters[randomIndex.Int64()])
		tmpLetters = tmpLetters[:randomIndex.Int64()] + tmpLetters[randomIndex.Int64()+1:]
	}

	return result, nil
}

func GenerateDecryptionKey(key map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range key {
		result[v] = k
	}
	return result
}
