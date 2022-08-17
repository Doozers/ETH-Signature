package main

import (
	"github.com/Doozers/ETH-Signature/ethsign"
)

func main() {
	ethsign.Sign("0x0", "Hello World")
	ethsign.Verify("Hello World", "0x0", "0xO", ethsign.TextHash)
}
