package tlv

import (
	"util/packet"
)

type Buffer struct {
	tlvVer int
	packet *packet.Packet
}

func (b *Buffer) WriteTlv(ver int) {
	b.tlvVer = ver
}

func (b *Buffer) WriteByte(i int) error {
	return b.packet.WriteByte(i)
}

func (b *Buffer) WriteBytes(bs []byte) error {
	return b.packet.WriteBytes(bs)
}

func (b *Buffer) WriteBoolean(z bool) error {
	return b.packet.WriteBoolean(z)
}

func (b *Buffer) WriteShort(i int) error {
	return b.packet.WriteShort(i)
}

func (b *Buffer) WriteInt(i int) error {
	return b.packet.WriteInt(i)
}

func (b *Buffer) WriteLong(i int) error {
	return b.packet.WriteLong(i)
}

func (b *Buffer) WriteString(s string) error {
	return b.packet.WriteString(s)
}

func (b *Buffer) ToByteArray() []byte {
	buffer := packet.CreateBuilder()
	_ = buffer.WriteShort(b.tlvVer)
	body := b.packet.Bytes()
	_ = buffer.WriteShort(len(body))
	_ = buffer.WriteBytes(body)
	return buffer.Bytes()
}
