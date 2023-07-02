package licensex

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	fmt.Println("-------------------------------License 生成-----------------------------------------")
	machineId, err := GetMachineId()
	if err != nil {
		panic(err)
	}

	li := &License{MachineId: machineId, Name: "test", Count: 0, Expiry: "2023-12-31"}
	data, err := li.ToBytes()
	if err != nil {
		panic(err)
	}

	publicKey, err := hex.DecodeString(licensePublicKey)
	if err != nil {
		panic(err)
	}

	ciphertext, err := RsaEncrypt(data, publicKey)
	if err != nil {
		panic(err)
	}
	encryptData := hex.EncodeToString(ciphertext)
	fmt.Println("公钥加密后的数据：", encryptData)
	nl, err := NewLicense(strings.NewReader(encryptData))
	if err != nil {
		panic(err)
	}

	fmt.Println(nl.Valid())

}
