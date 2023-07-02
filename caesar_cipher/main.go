// Author: Arkan M. Gerges<arkan.m.gerges@gmail.com>

package caesar_cipher

func Encrypt(offset int, plaintext string) (result string) {
	key := GenerateKey(offset)
	for _, c := range plaintext {
		if key[string(c)] == "" {
			result += string(c)
		} else {
			result += key[string(c)]
		}
	}

	return
}

func Decrypt(offset int, plaintext string) (result string) {
	key := GenerateDecryptionKeyBy(GenerateKey(offset))
	for _, c := range plaintext {
		if key[string(c)] == "" {
			result += string(c)
		} else {
			result += key[string(c)]
		}
	}
	return
}

func GenerateKey(offset int) (key map[string]string) {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+{}[]|\\:;\"',./<>?`~"
	lettersLen := len(letters)
	key = make(map[string]string)
	for i, c := range letters {
		key[string(c)] = string(letters[(i+offset)%lettersLen])
	}

	return
}

func GenerateDecryptionKeyBy(key map[string]string) (result map[string]string) {
	result = make(map[string]string)
	for _, v := range key {
		result[key[v]] = v
	}
	return
}
