package ethsign

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type sigProvider uint8

// sig provider
const (
	Metamask sigProvider = 1
	GoEth    sigProvider = 2
)

func Verify(originalMessage string, signature string, signerAddress string, sigProv sigProvider) (bool, error) {
	var recovered *ecdsa.PublicKey
	switch sigProv {
	case Metamask:
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
	case GoEth:
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
