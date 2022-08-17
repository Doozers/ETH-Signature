package ethsign

import (
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func Sign(privateKey string, message string) []byte {
	p, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	hash := crypto.Keccak256Hash([]byte(message))

	signature, err := crypto.Sign(hash.Bytes(), p)
	if err != nil {
		log.Fatal(err)
	}
	return signature
}
