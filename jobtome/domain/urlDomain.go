package domain

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"log"
	"math/big"
)

type Urls struct {
	OriginalUrl string `json:"original-url"`
	ShortUrl    string `json:"short-url"`
	Clicks      int
}

func (u *Urls) GenerateShortLink(initialLink string) string {
	urlHashBytes := sha256Of(initialLink)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return string(encoded)
}
