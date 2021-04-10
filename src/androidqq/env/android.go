package env

import (
	"util/cryptor/md5"
	"util/hex"
)

type Android struct {
	Imsi                 string
	AndroidId            string
	MachineName          string
	OsType               string
	MachineManufacturer  string
	MachineVersion       string
	MachineVersionNumber int
	WifiSSID             string
	WifiBSSID            string
	MacAddress           string
	Apn                  string
	ApnName              string

	Ksid []byte
}

func (a *Android) GetGuid() []byte {
	return md5.ToMd5BytesV2(a.AndroidId + a.MacAddress)
}

func (a *Android) GetKsid() []byte {
	if a.Ksid != nil {
		return a.Ksid
	}
	return hex.Str2Bytes("14751d8e7d633d9b06a392c357c675e5")
}

func (a *Android) SetKsid(ksid []byte) {
	a.Ksid = ksid
}

func (a *Android) GetTgtgKey() []byte {
	return md5.ToMd5Bytes(append(md5.ToMd5BytesV2(a.MacAddress), a.GetGuid()...))
}
