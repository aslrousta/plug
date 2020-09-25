package rand

import (
	secureRand "crypto/rand"
	"math/rand"
	"time"
)

const (
	alphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	alphabetSize = len(alphabet)
)

// String generates a random string consisting of letters and digits.
func String(size int) string {
	bytes := make([]byte, size)
	if _, err := secureRand.Read(bytes); err != nil {
		rand.Seed(time.Now().Unix())
		for index := range bytes {
			bytes[index] = byte(rand.Intn(alphabetSize))
		}
	}
	alphabet := []byte(alphabet)
	chars := make([]byte, size)
	for index := range bytes {
		chars[index] = alphabet[int(bytes[index])%alphabetSize]
	}
	return string(chars)
}
