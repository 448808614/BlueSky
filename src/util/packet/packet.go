package packet

import (
	"api"
	"bluesky"
	"util/bytes"
)

func MakeRequestPacket(cmd string, body, key, tgt []byte, seq int, status byte, info *api.ProtocolInfo) []byte {
	var tgtType = 0
	if tgt != nil {
		tgtType = 256
	}
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
	if info != nil {
		println(tgtType)
	} else {

	}

	return outBuilder.Bytes()
}
