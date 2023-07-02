package licensex

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/yansuan/go-zone-core/licensex/machine"
)

type License struct {
	MachineId string `json:"machineId"` //机器唯一ID
	Name      string `json:"name"`      //授权单位
	Count     int64  `json:"count"`     //授权数量
	Expiry    string `json:"expiry"`    //过期日期

	isValid bool //
}

func NewLicense(reader io.Reader) (*License, error) {
	// buf := new(bytes.Buffer)
	// _, err := buf.ReadFrom(reader)
	// if err != nil {
	// 	return nil, err
	// }

	// _, err := os.Stat(filename)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		machineId, err := GetMachineId()
	// 		if err != nil {
	// 			logger.Error("get license machine code error:", err)
	// 			return nil, err
	// 		}
	// 		fmt.Println(machineId)
	// 	}
	// 	return nil, err
	// }

	// f, err := os.Open(filename)
	// if err != nil {
	// 	machineId, err := GetMachineId()
	// 	if err != nil {
	// 		logger.Error("get license machine code error:", err)
	// 		return nil, err
	// 	}
	// 	fmt.Println(machineId)
	// 	return nil, err
	// }
	// defer f.Close()

	encyrptData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	bs, err := hex.DecodeString(string(encyrptData))
	if err != nil {
		return nil, err
	}

	privateKey, err := hex.DecodeString(licensePrivateKey)
	if err != nil {
		return nil, err
	}
	//decrypt
	data, err := RsaDecrypt(bs, []byte(privateKey))
	if err != nil {
		return nil, err
	}

	l := &License{}
	err = json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}
	l.isValid, err = l.valid()

	return l, nil
}

func (l *License) Valid() bool {
	return l.isValid
}

func (l *License) valid() (bool, error) {
	//machine
	machineId, err := GetMachineId()
	if err != nil {
		return false, err
	}

	if machineId != l.MachineId {
		return false, nil
	}

	if l.Expiry == "" {
		return true, nil
	}

	//expiry
	now := time.Now()
	expiry, err := time.Parse(time.DateOnly, l.Expiry)
	if err != nil {
		return false, err
	}
	if now.After(expiry) {
		return false, err
	}

	return true, nil
}

func (l *License) ToBytes() ([]byte, error) {
	return json.Marshal(l)
}

func GetMachineId() (string, error) {
	serialNumber, err := machine.GetSerialNumber()
	if err != nil {
		return "", err
	}
	uuid, err := machine.GetPlatformUUID()
	if err != nil {
		return "", err
	}

	cpuid, err := machine.GetCpuId()
	if err != nil {
		return "", err
	}

	macInfo, err := machine.GetMACAddress()
	if err != nil {
		return "", err
	}

	sign := Signature(serialNumber, uuid, cpuid, macInfo)
	return sign, nil
}
