package machine

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net"
	"runtime"
	"strings"
)

type MachineData struct {
	PlatformUUID string `json:"platformUUID"`
	SerialNumber string `json:"serialNumber"`
	CpuId        string `json:"cpuId"`
	Mac          string `json:"mac"`
}

func GetMachineData() (data MachineData) {
	if runtime.GOOS == "darwin" {
		sysInfo, _ := MacMachine{}.getMacSysInfo()
		sysInfo.Mac, _ = GetMACAddress()
		return sysInfo
	} else if runtime.GOOS == "linux" {
		machineData := MachineData{}
		machineData.SerialNumber, _ = LinuxMachine{}.getSerialNumber()
		machineData.PlatformUUID, _ = LinuxMachine{}.getPlatformUUID()
		machineData.CpuId, _ = LinuxMachine{}.getCpuId()
		machineData.Mac, _ = GetMACAddress()
		return machineData
	} else if runtime.GOOS == "windows" {
		machineData := MachineData{}
		machineData.SerialNumber, _ = WindowsMachine{}.getSerialNumber()
		machineData.PlatformUUID, _ = WindowsMachine{}.getPlatformUUID()
		machineData.CpuId, _ = WindowsMachine{}.getCpuId()
		machineData.Mac, _ = GetMACAddress()
		return machineData
	}
	return MachineData{}
}

func GetSerialNumber() (data string, err error) {
	if runtime.GOOS == "darwin" {
		return MacMachine{}.getSerialNumber()
	} else if runtime.GOOS == "linux" {
		return LinuxMachine{}.getSerialNumber()
	} else if runtime.GOOS == "windows" {
		return WindowsMachine{}.getSerialNumber()
	}
	return "", nil
}

func GetPlatformUUID() (data string, err error) {
	if runtime.GOOS == "darwin" {
		return MacMachine{}.getPlatformUUID()
	} else if runtime.GOOS == "linux" {
		return LinuxMachine{}.getPlatformUUID()
	} else if runtime.GOOS == "windows" {
		return WindowsMachine{}.getPlatformUUID()
	}
	return "", nil
}

func GetCpuId() (data string, err error) {
	if runtime.GOOS == "darwin" {
		return MacMachine{}.getCpuId()
	} else if runtime.GOOS == "linux" {
		return LinuxMachine{}.getCpuId()
	} else if runtime.GOOS == "windows" {
		return WindowsMachine{}.getCpuId()
	}
	return "", nil
}

func GetMACAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var mac string
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipnet, ok := address.(*net.IPNet)
				if ok && ipnet.IP.IsGlobalUnicast() {
					mac = netInterfaces[i].HardwareAddr.String()
					return mac, nil
				}
			}
		}
	}
	return mac, errors.New("无法获取到正确的MAC地址")
}

func GetMd5String(s string, upper bool, half bool) string {
	h := md5.New()
	h.Write([]byte(s))
	result := hex.EncodeToString(h.Sum(nil))
	if upper == true {
		result = strings.ToUpper(result)
	}
	if half == true {
		result = result[8:24]
	}
	return result
}
