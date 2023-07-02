package licensex

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRsa(t *testing.T) {
	prvKey, pubKey := GenRsaKey()

	fmt.Println("public key:")
	fmt.Println(hex.EncodeToString(pubKey))
	fmt.Println("private key:")
	fmt.Println(hex.EncodeToString(prvKey))
}
