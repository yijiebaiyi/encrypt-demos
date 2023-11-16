package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	keyStr := "123456789abcdefg"
	str := "helloworld"

	encodedStr := encode(str)
	fmt.Println("base64Encode: ", encodedStr)

	result, _ := encryptAES(encodedStr, keyStr, keyStr)
	fmt.Println("encryptAES :", result)
}

func encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func encryptAES(text, key, iv string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	padText := padPKCS7([]byte(text), block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, []byte(iv))

	cipherText := make([]byte, len(padText))
	mode.CryptBlocks(cipherText, padText)
	return hex.EncodeToString(cipherText), nil
}

func padPKCS7(input []byte, blockSize int) []byte {
	padding := blockSize - (len(input) % blockSize)
	padText := make([]byte, len(input)+padding)
	copy(padText, input)
	for i := len(input); i < len(padText); i++ {
		padText[i] = byte(padding)
	}
	return padText
}
