package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func hash256(text string) string {
	data := []byte(text)
	hash := sha256.Sum256(data)
	
	return fmt.Sprintf("%x", hash[:])
}

func hmac512(text string, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(text))
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
