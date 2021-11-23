package services

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func hashMD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func hash256SHA(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func hmacMD5(json []byte, apiPassword string) string {
	h := hmac.New(md5.New, []byte(apiPassword))
	h.Write(json)
	return hex.EncodeToString(h.Sum(nil))
}
