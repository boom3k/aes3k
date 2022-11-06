package aes3k

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Decrypt(data []byte, keyString string) []byte {
	// Byte array of the string

	// Key
	key := []byte(keyString)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(data) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := data[:aes.BlockSize]

	// Remove the IV from the data
	data = data[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from data
	stream.XORKeyStream(data, data)

	return data
}

func Encrypt(data []byte, keyString string) []byte {
	// Key
	key := []byte(keyString)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Empty array of 16 + data length
	// Include the IV at the beginning
	encryptedData := make([]byte, aes.BlockSize+len(data))

	// Slice of first 16 bytes
	iv := encryptedData[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from data to encryptedData
	stream.XORKeyStream(encryptedData[aes.BlockSize:], data)

	return encryptedData
}
