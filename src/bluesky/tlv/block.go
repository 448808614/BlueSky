package tlv

import (
	"bluesky/record"
	"github.com/gogo/protobuf/proto"
	"math/rand"
	login "protocol/protobuf"
	"strconv"
	"strings"
	"util/cryptor/md5"
	"util/cryptor/tea"
	"util/hex"
	"util/packet"
)

func (t *Tlv) T1() []byte {
	buffer := newBuffer(0x1)
	buffer.WriteShort(t.ProtocolInfo.IpVersion)
	buffer.WriteInt(0)
	buffer.WriteInt(t.BotAccount.Uin)
	buffer.WriteInt(t.Record.InitTime)
	buffer.WriteInt(0)
	buffer.WriteShort(0)
	return buffer.ToByteArray()
}

func (t *Tlv) T8() []byte {
	buffer := newBuffer(0x8)
	buffer.WriteShort(0)
	buffer.WriteInt(t.ProtocolInfo.LocalId)
	buffer.WriteShort(0)
	return buffer.ToByteArray()
}

func (t *Tlv) T18() []byte {
	buffer := newBuffer(0x18)
	buffer.WriteShort(t.ProtocolInfo.PingVersion)
	buffer.WriteInt(t.ProtocolInfo.SSoVersion)
	buffer.WriteInt(t.ProtocolInfo.SubAppId)
	buffer.WriteInt(0)
	buffer.WriteInt(t.BotAccount.Uin)
	buffer.WriteShort(0)
	buffer.WriteShort(0)
	return buffer.ToByteArray()
}

func (t *Tlv) T100() []byte {
	buffer := newBuffer(0x100)
	buffer.WriteShort(t.ProtocolInfo.DbVersion)
	buffer.WriteInt(t.ProtocolInfo.MsfSSoVersion)
	buffer.WriteInt(t.ProtocolInfo.SubAppId)
	buffer.WriteInt(t.ProtocolInfo.MsfAppId)
	buffer.WriteInt(0)
	buffer.WriteInt(0x21410e0)
	return buffer.ToByteArray()
}

func (t *Tlv) T104(dt104 []byte) []byte {
	buffer := newBuffer(0x104)
	buffer.WriteBytes(dt104)
	return buffer.ToByteArray()
}

func (t *Tlv) T107() []byte {
	buffer := newBuffer(0x107)
	buffer.WriteShort(0)
	buffer.WriteByte(0)
	buffer.WriteShort(0)
	buffer.WriteByte(1)
	return buffer.ToByteArray()
}

func (t *Tlv) T108(ksid []byte) []byte {
	buffer := newBuffer(0x108)
	if ksid == nil {
		ksid = hex.Str2Bytes("BF F3 F1 1C 63 EE 2C B1 7D 96 77 02 A3 6E 25 12")
	}
	buffer.WriteBytes(ksid)
	return buffer.ToByteArray()
}

func (t *Tlv) T116() []byte {
	buffer := newBuffer(0x116)
	buffer.WriteByte(0)
	// version
	buffer.WriteInt(t.ProtocolInfo.MiscBitmap)
	buffer.WriteInt(t.ProtocolInfo.SubSigMap)
	appIdArray := []int{0x5f5e10e2}
	buffer.WriteByte(len(appIdArray))
	for _, i := range appIdArray {
		buffer.WriteInt(i)
	}
	return buffer.ToByteArray()
}

func (t *Tlv) T1162() []byte {
	buffer := newBuffer(0x116)
	buffer.WriteByte(0)
	// version
	buffer.WriteInt(0x08F7FF7C)
	buffer.WriteInt(66560)
	appIdArray := []int{0x5f5e10e2}
	// login AppId --> 默认为16
	buffer.WriteByte(len(appIdArray))
	for _, i := range appIdArray {
		buffer.WriteInt(i)
	}
	return buffer.ToByteArray()
}

// T106 loginType 登录类型
// 1 为密码登录
func (t *Tlv) T106(loginType int) []byte {
	buffer := newBuffer(0x106)

	builder := packet.CreateBuilder()
	builder.WriteShort(t.ProtocolInfo.TgtVersion)
	builder.WriteInt(rand.Int())
	builder.WriteInt(t.ProtocolInfo.MsfSSoVersion)
	builder.WriteInt(t.ProtocolInfo.SubAppId)
	builder.WriteInt(0)
	builder.WriteInt(0)
	builder.WriteInt(t.BotAccount.Uin)
	builder.WriteInt(t.Record.InitTime)
	builder.WriteInt(0)
	// 上面这个东西是IP傲
	builder.WriteByte(1)
	builder.WriteBytes(t.BotAccount.Md5Password())
	builder.WriteBytes(t.Android.GetTgtgKey())
	builder.WriteInt(0)
	builder.WriteBoolean(true)
	builder.WriteBytes(t.Android.GetGuid())
	builder.WriteInt(t.ProtocolInfo.MsfAppId)
	builder.WriteInt(loginType)

	user := strconv.Itoa(t.BotAccount.Uin)
	builder.WriteShort(len(user))
	builder.WriteString(user)

	builder.WriteShort(0)

	buffer.WriteBytes(tea.NewCipher(t.BotAccount.Md5UinPassword()).Encrypt(builder.Bytes()))

	return buffer.ToByteArray()
}

func (t *Tlv) T124() []byte {
	buffer := newBuffer(0x124)

	buffer.WriteStringWithShortSize("android")
	buffer.WriteStringWithShortSize(t.Android.MachineVersion)
	switch t.Android.Apn {
	case "5g", "wap", "net", "4gnet", "3gwap", "cncc", "cmcc", "3gnet", "cmwap", "uniwap":
		buffer.WriteShort(1)
	case "wifi":
		buffer.WriteShort(2)
	default:
		buffer.WriteShort(0)
	}
	// 在线状态书写 by: QQ逆向
	buffer.WriteStringWithShortSize(t.Android.Apn)
	buffer.WriteStringWithSize(t.Android.ApnName)
	return buffer.ToByteArray()
}

func (t *Tlv) T128() []byte {
	buffer := newBuffer(0x128)
	buffer.WriteShort(0)
	buffer.WriteBoolean(true)
	buffer.WriteBoolean(false)
	buffer.WriteBoolean(false)
	buffer.WriteInt(0x01000000)
	buffer.WriteStringWithShortSize(t.Android.MachineName)
	buffer.WriteBytesWithShortSize(t.Android.GetGuid())
	buffer.WriteStringWithShortSize(t.Android.MachineManufacturer)
	return buffer.ToByteArray()
}

func (t *Tlv) T141() []byte {
	buffer := newBuffer(0x141)
	buffer.WriteShort(1)
	buffer.WriteStringWithShortSize(t.Android.ApnName)
	switch t.Android.Apn {
	case "5g", "wap", "net", "4gnet", "3gwap", "cncc", "cmcc", "3gnet", "cmwap", "uniwap":
		buffer.WriteShort(1)
	case "wifi":
		buffer.WriteShort(2)
	default:
		buffer.WriteShort(0)
	}
	buffer.WriteStringWithShortSize(t.Android.Apn)
	return buffer.ToByteArray()
}

func (t *Tlv) T142() []byte {
	buffer := newBuffer(0x142)
	buffer.WriteShort(0)
	buffer.WriteStringWithShortSize(t.ProtocolInfo.PackageName)
	return buffer.ToByteArray()
}

func (t *Tlv) T109() []byte {
	buffer := newBuffer(0x109)
	buffer.WriteBytes(md5.StrToMd5Bytes(t.Android.AndroidId))
	return buffer.ToByteArray()
}

func (t *Tlv) T144() []byte {
	buffer := newBuffer(0x144)
	builder := packet.CreateBuilder()
	tlvArray := []int{
		0x109, 0x52d, 0x124, 0x128, 0x16e,
	}
	builder.WriteShort(len(tlvArray))
	for _, ver := range tlvArray {
		switch ver {
		case 0x109:
			builder.WriteBytes(t.T109())
		case 0x52d:
			builder.WriteBytes(t.T52d())
		case 0x124:
			builder.WriteBytes(t.T124())
		case 0x128:
			builder.WriteBytes(t.T128())
		case 0x16e:
			builder.WriteBytes(t.T16e())
		}
	}
	buffer.WriteBytes(tea.NewCipher(t.Android.GetTgtgKey()).Encrypt(builder.Bytes()))
	return buffer.ToByteArray()
}

func (t *Tlv) T52d() []byte {
	buffer := newBuffer(0x52d)
	deviceReport := login.DeviceReport{
		Bootloader:  []byte("unknown"),
		Version:     []byte("Linux version 4.19.113-perf-gb3dd08fa2aaa (builder@c5-miui-ota-bd143.bj) (clang version 8.0.12 for Android NDK) #1 SMP PREEMPT Thu Feb 4 04:37:10 CST 2021;"),
		Codename:    []byte("REL"),
		Incremental: []byte("20.8.13"),
		Fingerprint: []byte("Xiaomi/vangogh/vangogh:11/RKQ1.200826.002/21.2.4:user/release-keys"),
		BootId:      []byte(""),
		AndroidId:   []byte(t.Android.AndroidId),
		Baseband:    []byte(""),
		InnerVer:    []byte("21.2.4"),
	}
	bytes, _ := proto.Marshal(&deviceReport)
	buffer.WriteBytes(bytes)
	return buffer.ToByteArray()
}

func (t *Tlv) T16e() []byte {
	buffer := newBuffer(0x16e)
	buffer.WriteString(t.Android.MachineName)
	return buffer.ToByteArray()
}

func (t *Tlv) T145() []byte {
	buffer := newBuffer(0x145)
	buffer.WriteBytes(t.Android.GetGuid())
	return buffer.ToByteArray()
}

func (t *Tlv) T147() []byte {
	buffer := newBuffer(0x147)
	buffer.WriteInt(t.ProtocolInfo.SubAppId)
	buffer.WriteStringWithShortSize(t.ProtocolInfo.PackageVersion)
	buffer.WriteBytesWithShortSize(hex.Str2Bytes(t.ProtocolInfo.MsfSdkMd5))
	return buffer.ToByteArray()
}

func (t *Tlv) T154(seq int) []byte {
	buffer := newBuffer(0x154)
	buffer.WriteInt(seq)
	return buffer.ToByteArray()
}

func (t *Tlv) T16b() []byte {
	buffer := newBuffer(0x16b)
	buffer.WriteShort(len(domains))
	for _, domain := range domains {
		buffer.WriteStringWithShortSize(domain)
	}
	return buffer.ToByteArray()
}

// T16a 刷新SKey
func (t *Tlv) T16a() []byte {
	buffer := newBuffer(0x16a)
	// 填入参数noPicSig
	// noPicSig 由腾讯服务器提供
	buffer.WriteBytes(t.Record.GetKey(record.NoPicSig).Source)
	return buffer.ToByteArray()
}

func (t *Tlv) T174() []byte {
	buffer := newBuffer(0x174)
	buffer.WriteBytes(t.Record.GetKey(record.TLV174).Source)
	return buffer.ToByteArray()
}

func (t *Tlv) T177() []byte {
	buffer := newBuffer(0x177)
	buffer.WriteBoolean(true)
	buffer.WriteInt(t.ProtocolInfo.BuildTime)
	buffer.WriteStringWithShortSize(t.ProtocolInfo.BuildVersion)
	return buffer.ToByteArray()
}

// T17a 刷新短信验证码
func (t *Tlv) T17a() []byte {
	buffer := newBuffer(0x17a)
	// QQ内写死为9 无法改变 不知道是什么东西
	buffer.WriteInt(9)
	return buffer.ToByteArray()
}

func (t *Tlv) T17c(code string) []byte {
	buffer := newBuffer(0x17c)
	buffer.WriteStringWithShortSize(code)
	return buffer.ToByteArray()
}

func (t *Tlv) T187() []byte {
	buffer := newBuffer(0x187)
	buffer.WriteBytes(md5.StrToMd5Bytes(t.Android.MacAddress))
	return buffer.ToByteArray()
}

func (t *Tlv) T188() []byte {
	buffer := newBuffer(0x188)
	buffer.WriteBytes(md5.StrToMd5Bytes(t.Android.AndroidId))
	return buffer.ToByteArray()
}

func (t *Tlv) T191() []byte {
	buffer := newBuffer(0x191)
	buffer.WriteByte(0x82)
	// 82为滑块验证 0 为自动 1 为字母（协议不支持）
	return buffer.ToByteArray()
}

func (t *Tlv) T193(ticket string) []byte {
	buffer := newBuffer(0x193)
	buffer.WriteString(ticket)
	return buffer.ToByteArray()
}

func (t *Tlv) T194() []byte {
	buffer := newBuffer(0x194)
	buffer.WriteBytes(md5.StrToMd5Bytes(t.Android.Imsi))
	return buffer.ToByteArray()
}

func (t *Tlv) T197() []byte {
	buffer := newBuffer(0x197)
	// DeviceLockMobileType
	// 初始化为 0
	buffer.WriteByte(0)
	return buffer.ToByteArray()
}

func (t *Tlv) T198() []byte {
	buffer := newBuffer(0x198)
	buffer.WriteByte(0)
	return buffer.ToByteArray()
}

func (t *Tlv) T202() []byte {
	buffer := newBuffer(0x202)
	buffer.WriteBytesWithShortSize(md5.StrToMd5Bytes(t.Android.WifiBSSID))
	buffer.WriteStringWithShortSize("\"" + t.Android.WifiSSID + "\"")
	return buffer.ToByteArray()
}

// TODO("TLV-400没有开始写")

func (t *Tlv) T401() []byte {
	buffer := newBuffer(0x401)
	builder := packet.CreateBuilder()
	builder.WriteBytes(t.Android.GetGuid())
	builder.WriteString("1234567890123456")
	builder.WriteBytes(t.Record.GetKey(record.TLV402).Source)
	buffer.WriteBytes(md5.BsToMd5Bytes(builder.Bytes()))
	return buffer.ToByteArray()
}

func (t *Tlv) T402() []byte {
	buffer := newBuffer(0x402)
	buffer.WriteBytes(t.Record.GetKey(record.TLV402).Source)
	return buffer.ToByteArray()
}

func (t *Tlv) T403() []byte {
	buffer := newBuffer(0x403)
	buffer.WriteBytes(t.Record.GetKey(record.TLV403).Source)
	return buffer.ToByteArray()
}

func (t *Tlv) T511() []byte {
	buffer := newBuffer(0x511)
	buffer.WriteShort(len(domains))
	for _, domain := range domains {
		start := strings.Index(domain, "(")
		end := strings.Index(domain, ")")
		var b byte = 1
		if start == 0 || end > 0 {
			i, err := strconv.Atoi(domain[start+1 : end])
			if err == nil {
				if (1048576 & i) <= 0 {
					b = 0
				}
				if i&0x08000000 > 0 {
					b = b | 2
				}
			}
		}
		buffer.WriteByte(int(b))
		buffer.WriteStringWithShortSize(domain)
	}
	return buffer.ToByteArray()
}

func (t *Tlv) T516() []byte {
	buffer := newBuffer(0x516)
	buffer.WriteInt(0)
	return buffer.ToByteArray()
}

func (t *Tlv) T521() []byte {
	buffer := newBuffer(0x521)
	buffer.WriteInt(0)
	buffer.WriteShort(0)
	return buffer.ToByteArray()
}

func (t *Tlv) T525(flag int) []byte {
	buffer := newBuffer(0x525)
	buffer.WriteShort(1)
	// 1 普通登录
	// 2 假锁登录
	if flag == 1 {
		buffer.WriteBytes(t.T536(false))
	} else {
		buffer.WriteBytes(t.T536(true))
	}
	return buffer.ToByteArray()
}

func (t *Tlv) T536(lock bool) []byte {
	buffer := newBuffer(0x536)
	buffer.WriteByte(1)
	if lock {
		buffer.WriteByte(3) // cNum
		buffer.WriteBytes(hex.Str2Bytes("00 00 00 00 6F E2 BD 2E 04 DF 68 5C 58 5F 8A F7 FE 20 02 F8 A1 00 00 00 00 99 68 42 96 04 DF 68 5C 58 5F 8A FD 41 20 02 F8 A1 00 00 00 00 62 5A BF 50 04 DF 68 5C 58 5F 8A FD 42 20 02 F8 A1"))
	} else {
		buffer.WriteByte(0) // cNum
	}
	// 第一次登录没有LoginExtraData
	// 所以说没有下文
	// 01 -- 固定为 1
	// 03 -- cNum 代表有3个LoginExtraData
	// 00 00 00 00 6F E2 BD 2E 04 DF 68 5C 58 5F 8A F7 FE 20 02 F8 A1 00 00 00 00 99 68 42 96 04 DF 68 5C 58 5F 8A FD 41 20 02 F8 A1 00 00 00 00 62 5A BF 50 04 DF 68 5C 58 5F 8A FD 42 20 02 F8 A1
	return buffer.ToByteArray()
}

func (t *Tlv) T542() []byte {
	buffer := newBuffer(0x542)
	buffer.WriteByte(0)
	buffer.WriteByte(0)
	return buffer.ToByteArray()
}

func (t *Tlv) T544() []byte {
	buffer := newBuffer(0x544)
	buffer.WriteInt(2009)
	// 不知道有什么用
	buffer.WriteInt(0)

	randSize := 32
	requestInfoSize := 8

	buffer.WriteShort(2 + randSize + requestInfoSize)
	// 包括RandSize长度的长度所以加2
	buffer.WriteShort(randSize)
	randBs := packet.RandBytes(randSize)
	buffer.WriteBytes(randBs)

	buffer.WriteShort(requestInfoSize) // 接下来的东西的长度
	buffer.WriteInt(0)
	buffer.WriteInt(int(rand.Uint32()))
	buffer.WriteShort(3)

	buffer.WriteBoolean(true)
	buffer.WriteInt(4)
	buffer.WriteShort(0)
	// 0 为 2bit
	// 1 为 4bit --> buffer.WriteInt(1)

	buffer.WriteBytes(func() []byte {
		builder := packet.CreateBuilder()
		builder.WriteByte(0)
		builder.WriteByte(0)
		builder.WriteByte(0)

		// 看不懂随机吧
		key := packet.RandBytes(16)
		if true {
			key[0] = 1
		}
		return builder.Bytes()
	}())

	buffer.WriteStringWithShortSize(t.ProtocolInfo.PackageName)

	buffer.WriteString("A6B745BF24A2C277527716F6F36EB68D")
	// 固定了，太奇怪了

	buffer.WriteBytes(packet.RandBytes(4))

	buffer.WriteInt(0)

	return buffer.ToByteArray()
}

func newBuffer(ver int) *Buffer {
	buffer := Buffer{
		tlvVer: ver,
		packet: packet.CreateBuilder(),
	}
	return &buffer
}
