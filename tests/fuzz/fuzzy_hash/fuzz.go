package fuzz

import "github.com/caffix/jarm-go"

// Fuzz uses go-fuzz to test RawHashToFuzzyHash()
func Fuzz(data []byte) int {
	jarm.RawHashToFuzzyHash(string(data))
	return 1
}
