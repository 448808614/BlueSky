package tlv

import (
	"androidqq/account"
	"androidqq/env"
	"androidqq/record"
	"api"
	"github.com/gogo/protobuf/proto"
	"math/rand"
	login "protocol/protobuf"
	"strconv"
	"util/cryptor/md5"
	"util/cryptor/tea"
	"util/hex"
	"util/packet"
)

type Tlv struct {
	BotAccount   *account.BotAccount
	ProtocolInfo *api.ProtocolInfo
	Record       *record.BotRecord
	Android      *env.Android
}

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

func (t *Tlv) T109() []byte {
	buffer := newBuffer(0x109)
	buffer.WriteBytes(md5.ToMd5BytesV2(t.Android.AndroidId))
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

/**
fun t144(): ByteArray {
val tlvBuilder = TlvBuilder(0x144)
val builder = ByteBuilder()
builder.writeShort(5)
builder.writeBytes(t109())
builder.writeBytes(t52d())
builder.writeBytes(t124())
builder.writeBytes(t128())
builder.writeBytes(t16e())
val data = TeaUtil.encrypt(builder.toByteArray(), Android.getTgtgKey())
tlvBuilder.writeBytes(data)
return tlvBuilder.toByteArray()
}

fun t145(): ByteArray = TlvBuilder(0x145)
.writeBytes(Android.getGuid())
.toByteArray()

fun t147(): ByteArray = TlvBuilder(0x147)
.writeInt(iqq.subAppId())
.writeStringWithShortSize(iqq.packageVersion())
.writeBytesWithShortSize(iqq.tencentSdkMd5())
.toByteArray()

fun t154(): ByteArray = TlvBuilder(0x154)
.writeInt(seq)
.toByteArray()

fun t16a(): ByteArray = TlvBuilder(0x16a)
.writeHex("20 B5 33 79 18 79 9C AB E4 4A 8E F8 0D 66 84 B81F 8C 15 24 AD 46 D6 D7 7A AF 24 6A 09 16 0A 59AF 22 ED 5B 14 A8 B4 78 36 F2 AC 9A 34 61 15 3A")
.toByteArray()

 * 登录过 设备所显示名称

t16b(): ByteArray {
val tlvBuilder = TlvBuilder(0x16b)
val domains = arrayOf(
"tenpay.com",
"qzone.qq.com",
"qun.qq.com",
"mail.qq.com",
"openmobile.qq.com",
"qzone.com",
"game.qq.com",
"vip.qq.com"
)
tlvBuilder.writeShort(domains.size)
for (s in domains) {
tlvBuilder.writeBytesWithShortSize(s.toByteArray())
}
return tlvBuilder.toByteArray()
}

private fun t16e(): ByteArray {
val tlvBuilder = TlvBuilder(0x16e)
tlvBuilder.writeString(Android.machineName)
return tlvBuilder.toByteArray()
}

fun t174(dt174: ByteArray?): ByteArray {
val tlvBuilder = TlvBuilder(0x174)
tlvBuilder.writeBytes(dt174)
return tlvBuilder.toByteArray()
}

fun t177(): ByteArray {
val tlvBuilder = TlvBuilder(0x177)
tlvBuilder.writeBoolean(true)
tlvBuilder.writeInt(iqq.buildTime())
tlvBuilder.writeStringWithShortSize(iqq.buildVersion())
return tlvBuilder.toByteArray()
}

fun t17a(i: Long): ByteArray {
val tlvBuilder = TlvBuilder(0x17a)
tlvBuilder.writeInt(i)
return tlvBuilder.toByteArray()
}

fun t17c(code: String): ByteArray {
val tlvBuilder = TlvBuilder(0x17c)
tlvBuilder.writeShort(code.length)
tlvBuilder.writeString(code)
return tlvBuilder.toByteArray()
}

fun t187(): ByteArray {
val tlvBuilder = TlvBuilder(0x187)
tlvBuilder.writeBytes(MD5.toMD5Byte(Android.macAddress))
return tlvBuilder.toByteArray()
}

fun t188(): ByteArray {
val tlvBuilder = TlvBuilder(0x188)
tlvBuilder.writeBytes(MD5.toMD5Byte(Android.androidId))
return tlvBuilder.toByteArray()
}

fun t191(): ByteArray {
val tlvBuilder = TlvBuilder(0x191)
tlvBuilder.writeByte(0x82)
return tlvBuilder.toByteArray()
}

fun t194(): ByteArray {
val tlvBuilder = TlvBuilder(0x194)
tlvBuilder.writeBytes(MD5.toMD5Byte(Android.imsi))
return tlvBuilder.toByteArray()
}

fun t197(): ByteArray {
val tlvBuilder = TlvBuilder(0x197)
tlvBuilder.writeByte(0.toByte())
return tlvBuilder.toByteArray()
}

fun t198(): ByteArray {
val tlvBuilder = TlvBuilder(0x198)
tlvBuilder.writeByte(0.toByte())
return tlvBuilder.toByteArray()
}

fun t202(): ByteArray {
val tlvBuilder = TlvBuilder(0x202)
tlvBuilder.writeBytesWithShortSize(MD5.toMD5Byte(Android.wifiBSSID))
tlvBuilder.writeStringWithShortSize("\"" + Android.wifiSSID + "\"")
return tlvBuilder.toByteArray()
}

fun t400(): ByteArray {
val tlvBuilder = TlvBuilder(0x400)
tlvBuilder.writeHex("D1387BC477015873D624BB495618F37A3096BCB21757E66741E1E5E090E6DD293C402D0003B169879C5B95BB5A21028062CD406335AFE249A508144C26A18A42B3FF12D1A1EB95E8")
return tlvBuilder.toByteArray()
}

fun t401(dt402: ByteArray?): ByteArray {
val tlvBuilder = TlvBuilder(0x401)
val builder = ByteBuilder()
builder.writeBytes(Android.getGuid())
builder.writeBytes(QQUtil.get_mpasswd().toByteArray())
builder.writeBytes(dt402)
tlvBuilder.writeBytes(MD5.toMD5Byte(builder.toByteArray()))
return tlvBuilder.toByteArray()
}

fun t402(dt402: ByteArray?): ByteArray {
val tlvBuilder = TlvBuilder(0x402)
tlvBuilder.writeBytes(dt402)
return tlvBuilder.toByteArray()
}

fun t403(dt403: ByteArray?): ByteArray {
val tlvBuilder = TlvBuilder(0x403)
tlvBuilder.writeBytes(dt403)
return tlvBuilder.toByteArray()
}

fun t511(): ByteArray {
val tlvBuilder = TlvBuilder(0x511)
val domains = arrayOf(
"office.qq.com",
"qun.qq.com",
"gamecenter.qq.com",
"docs.qq.com",
"mail.qq.com",
"ti.qq.com",
"vip.qq.com",
"tenpay.qq.com",
"qqqweb.qq.com",
"qzone.qq.com",
"mma.qq.com",
"game.qq.com",
"openmobile.qq.com",
"conect.qq.com" // "y.qq.com",
// "v.qq.com"
)
tlvBuilder.writeShort(domains.size)
for (domain in domains) {
val start = domain.indexOf('(')
val end = domain.indexOf(')')
var b: Byte = 1
if (start == 0 || end > 0) {
val i = domain.substring(start + 1, end).toInt()
b = if (1048576 and i > 0) {
1
} else {
0
}
if (i and 0x08000000 > 0) {
b = (b or 2.toByte())
}
}
tlvBuilder.writeByte(b)
tlvBuilder.writeStringWithShortSize(domain)
}
return tlvBuilder.toByteArray()
}

fun t516(): ByteArray {
val tlvBuilder = TlvBuilder(0x516)
tlvBuilder.writeInt(0)
return tlvBuilder.toByteArray()
}

fun t521(): ByteArray {
val tlvBuilder = TlvBuilder(0x521)
tlvBuilder.writeInt(0)
tlvBuilder.writeShort(0)
return tlvBuilder.toByteArray()
}

fun t525(): ByteArray {
val tlvBuilder = TlvBuilder(0x525)
tlvBuilder.writeShort(1)
tlvBuilder.writeBytes(t536())
return tlvBuilder.toByteArray()
}

private fun t536(): ByteArray {
val tlvBuilder = TlvBuilder(0x536)
tlvBuilder.writeHex(
"""
01 03 00 00 00 00 6F E2
BD 2E 04 DF 68 5C 58 5F
8A F7 FE 20 02 F8 A1 00
00 00 00 99 68 42 96 04
DF 68 5C 58 5F 8A FD 41
20 02 F8 A1 00 00 00 00
62 5A BF 50 04 DF 68 5C
58 5F 8A FD 42 20 02 F8
A1
""".trimIndent()
)
return tlvBuilder.toByteArray()
}


fun t542(): ByteArray {
val tlvBuilder = TlvBuilder(0x542)
tlvBuilder.writeByte(0.toByte())
tlvBuilder.writeByte(0.toByte())
return tlvBuilder.toByteArray()
}

fun t544(): ByteArray {
val tlvBuilder = TlvBuilder(0x544)
tlvBuilder.writeHex(
"""
00 00 07 D9 00 00 00 00
00 2E 00 20 15 97 BF B2
50 07 9C 86 AF 7A FB 53
64 4F 39 97 E9 0A 15 91
83 AD F1 20 CC 89 F8 75
28 63 5C 3E 00 08 00 00
00 00 00 00 50 C9 00 03
01 00 00 00 04 00 00 00
03 00 00 00 01 75 37 33
F4 38 29 75 6F 47 67 32
76 48 33 70 00 14 63 6F
6D 2E 74 65 6E 63 65 6E
74 2E 6D 6F 62 69 6C 65
71 71 41 36 42 37 34 35
42 46 32 34 41 32 43 32
37 37 35 32 37 37 31 36
46 36 46 33 36 45 42 36
38 44 05 E7 AD 8C 00 00
00 00 00 00 02 00 00 01
00 10 D8 44 E7 DC BB E2
17 CB 9C 77 7F B0 FF B7
B7 42
""".trimIndent()
)
return tlvBuilder.toByteArray()
}
*/

func newBuffer(ver int) *Buffer {
	buffer := Buffer{
		tlvVer: ver,
		packet: packet.CreateBuilder(),
	}
	return &buffer
}
