package main

import (
	"crypto/rand"
	"encoding/base64"
    "log"
)

func CreateRandomString(length int) (string) {
	bytes := make([]byte, length)
	_,err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}