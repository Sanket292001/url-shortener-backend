package utilities

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
	"net/url"
)

func ValidUrl(Url string) (bool, error) {
	_, err := url.ParseRequestURI(Url)
	return err == nil, err
}

func RandomToken(length int) string {
	randomBytes := make([]byte, 32)
    _, err := rand.Read(randomBytes)
    if err != nil {
        panic(err)
    }
    return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

func GenerateHash(data string) [16]byte {
	return md5.Sum([]byte(data))
}

func GenerateHashedString(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}