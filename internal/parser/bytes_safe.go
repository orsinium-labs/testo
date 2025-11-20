package parser

import (
	"strconv"
)

func equalStr(b *[]byte, s string) bool {
	return string(*b) == s
}

func parseFloat(b *[]byte) (float64, error) {
	return strconv.ParseFloat(string(*b), 64)
}

func bytesToString(b *[]byte) string {
	return string(*b)
}

func stringToBytes(s string) []byte {
	return []byte(s)
}
