package main

import (
	"fmt"

	"github.com/Doozers/ETH-Signature/ethsign"
)

func main() {
	fmt.Println(ethsign.VerifyEthSignature("Hello World", "0x00d5f13a00b060106d5a001ede826ebd00b1e0c468bce006b6ac05b98d7077aa2e30c6fcf5febebbd06736626c7d7c52ee39ff45494adb5526174dfb58ad35d31c", "0x6192b1554434f6530ab412831A1DA2Fe8777E23c"))
}
