package core

import (
	"crypto/hmac"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
)

func GenerateToken(secret string, nonce []byte) string {
	hmac := hmac.New(sha512.New, []byte(secret))
	hmac.Write(nonce)
	digest := hmac.Sum(nil)

	result := make([]byte, 0, len(digest)+len(nonce))
	result = append(result, nonce...)
	result = append(result, digest...)
	return base64.StdEncoding.EncodeToString(result)
}

func VerifyToken(secret string, token string) bool {
	tokenData, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false
	}

	hmac := hmac.New(sha512.New, []byte(secret))
	if len(tokenData) <= hmac.Size() {
		return false
	}

	nonce := tokenData[0 : len(tokenData)-hmac.Size()]
	digest := tokenData[len(tokenData)-hmac.Size():]
	hmac.Write(nonce)
	computedDigest := hmac.Sum(nil)

	return subtle.ConstantTimeCompare(digest, computedDigest) == 1
}
