package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Check(content, encrypted string) bool {
	return strings.EqualFold(Md5Encode(content), encrypted)
}

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
