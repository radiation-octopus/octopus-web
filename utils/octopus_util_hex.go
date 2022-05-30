package utils

import "encoding/hex"

// Encode将b编码为前缀为0x的十六进制字符串。
func Encode(b []byte) string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}
