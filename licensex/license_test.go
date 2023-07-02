package licensex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestLicense(t *testing.T) {
	license1 := &License{}
	license1.MachineId = "9f6663a42a19333e7f52703ed8a5b3734bd8348e"
	license1.Expiry = "2023-06-09"
	license1.Name = "company1"
	license1.Count = 10

	b, err := json.Marshal(license1)
	if err != nil {
		panic(err)
	}

	ciphertext, err := RsaEncrypt(b, []byte(licensePublicKey))
	if err != nil {
		panic(err)
	}

	l, err := NewLicense(bytes.NewBuffer(ciphertext))
	if err != nil {
		panic(err)
	}
	fmt.Println(l.Name)
	fmt.Println(l.isValid)
	// if err != nil {
	// 	machineId, err := GetMachineId()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("License Code:", machineId)
	// 	return
	// }

}
