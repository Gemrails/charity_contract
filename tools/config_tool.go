package tools

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha8(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs[:4])
}

func Skey(s string) string {
	return ""
}
