package packet

import (
	"bluesky"
	"util/bytes"
	"util/cryptor/tea"
)

func MakeRequestPacket(cmd string, data, key, tgt, ksid []byte, seq int, status byte, msfAppId int, protocolVersion, androidId *string, needSize bool) []byte {
	var tgtType = 0
	if tgt != nil {
		tgtType = 256
	} else {
		tgt = []byte{}
	}
	builder := bytes.CreateBuilder()
	outBuilder := bytes.CreateBuilder()
	switch status {
	case bluesky.HasLogin:
		outBuilder.WriteInt(0xA)
		outBuilder.WriteByte(1)
	case bluesky.NoLogin:
		outBuilder.WriteInt(0xA)
		outBuilder.WriteByte(2)
	case bluesky.Online:
		outBuilder.WriteInt(0xB)
		outBuilder.WriteByte(1)
		outBuilder.WriteInt(seq)
	default:
		panic("错误状态码，无法进行组包...")
	}
	outBuilder.WriteByte(0)
	bodyBuilder := bytes.CreateBuilder()
	headBuilder := bytes.CreateBuilder()
	if androidId != nil {
		headBuilder.WriteInt(seq)
		headBuilder.WriteInt(msfAppId)
		headBuilder.WriteInt(msfAppId)
		headBuilder.WriteInt(16777216)
		headBuilder.WriteInt(0)
		headBuilder.WriteInt(tgtType)
		headBuilder.WriteBytesWithSize(tgt, 4)
		headBuilder.WriteStringWithSize(cmd, 4)
		headBuilder.WriteBytesWithSize(bytes.RandBytes(4), 4)
		headBuilder.WriteStringWithSize(*androidId, 4)
		if ksid == nil {
			ksid = []byte{}
		}
		headBuilder.WriteBytesWithSize(ksid, 4)
		headBuilder.WriteStringWithShortSize(*protocolVersion, 2)
		// head.writeInt(4);
		// 跳过ECDH
	} else {
		headBuilder.WriteStringWithSize(cmd, 4)
		headBuilder.WriteBytesWithSize(bytes.RandBytes(4), 4)
		headBuilder.WriteInt(4)
	}
	head := headBuilder.Bytes()
	bodyBuilder.WriteInt(len(head) + 4)
	bodyBuilder.WriteBytes(head)
	if needSize {
		bodyBuilder.WriteInt(len(data) + 4)
	}
	bodyBuilder.WriteBytes(data)
	eData := tea.NewCipher(key).Encrypt(bodyBuilder.Bytes())
	outBuilder.WriteBytes(eData)
	body := outBuilder.Bytes()
	builder.WriteInt(len(body) + 4)
	builder.WriteBytes(body)
	return builder.Bytes()
}
