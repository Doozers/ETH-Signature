package ethsign

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifyEthSignature(originalMessage string, signature string, signerAddress string) (bool, error) {
	hash := accounts.TextHash([]byte(originalMessage))

	signatureBytes, err := hexutil.Decode(signature)
	if err != nil {
		return false, err
	}
	signatureBytes[crypto.RecoveryIDOffset] -= 27

	recovered, err := crypto.SigToPub(hash, signatureBytes)
	if err != nil {
		return false, err
	}

	return crypto.PubkeyToAddress(*recovered).Hex() == signerAddress, nil
}
