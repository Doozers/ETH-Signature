package ethsign

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type hash uint8

const (
	TextHash  hash = 1
	Keccak256 hash = 2
)

func Verify(originalMessage string, signature string, signerAddress string, algoHash hash) (bool, error) {
	var recovered *ecdsa.PublicKey
	switch algoHash {
	case TextHash:
		hashByte := accounts.TextHash([]byte(originalMessage))
		signatureBytes, err := hexutil.Decode(signature)
		if err != nil {
			return false, err
		}

		signatureBytes[crypto.RecoveryIDOffset] -= 27

		recovered, err = crypto.SigToPub(hashByte, signatureBytes)
		if err != nil {
			return false, err
		}
	case Keccak256:
		hash := crypto.Keccak256Hash([]byte(originalMessage))
		signatureBytes, err := hexutil.Decode(signature)
		if err != nil {
			return false, err
		}
		recovered, err = crypto.SigToPub(hash.Bytes(), signatureBytes)
		if err != nil {
			return false, err
		}
	}
	return crypto.PubkeyToAddress(*recovered).Hex() == signerAddress, nil
}
