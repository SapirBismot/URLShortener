package url_shortener

import (
	"fmt"
	"crypto/sha256"
	"github.com/alextanhongpin/base62"
	"math/big"
)

func createUrl(longUrl string) string {
	hashUrl := sha256.Sum256([]byte(longUrl)])
	bytesNumber := new(big.Int).SetBytes(hashUrl[:]).Uint64()
	shortUrl := base62.Encode(bytesNumber)
	return shortUrl
}