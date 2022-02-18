package security

import (
    "encoding/base64"
	"crypto/sha256"
	"encoding/hex"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/md5"
    "strings"
    "io"
)

func MD5(input string) string {
	hashBinary := md5.Sum([]byte(strings.ToUpper(input)))
    return hex.EncodeToString(hashBinary[:])
}

func SH1(input string) string {
	hash := sha1.New()
	io.WriteString(hash, input)

	return hex.EncodeToString(hash.Sum(nil))
}

func HmacSHA256(input string, key string) string {
	sig := hmac.New(sha256.New, []byte(key))
	sig.Write([]byte(input))

	return Base64url(sig.Sum(nil))
}

func Base64url(bytes []byte) string {
	return base64.RawURLEncoding.EncodeToString(bytes)
}