package gopie

import (
	//nolint:gosec
	"crypto/md5"
	"encoding/hex"
)

// MD5 simply calculates md5 hash of string or []byte.
func MD5[T string | []byte](input T) string {
	//nolint:gosec
	hash := md5.New()
	hash.Write([]byte(input))
	hashInBytes := hash.Sum(nil)

	return hex.EncodeToString(hashInBytes)
}
