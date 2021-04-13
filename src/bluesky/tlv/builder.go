package tlv

import (
	"util/bytes"
)

type Buffer struct {
	tlvVer int
	packet *bytes.Builder
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

func (b *Buffer) WriteStringWithShortSize(s string, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	return b.packet.WriteStringWithShortSize(s, a)
}

func (b *Buffer) WriteStringWithSize(s string, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	return b.packet.WriteStringWithSize(s, a)
}

func (b *Buffer) WriteBytesWithShortSize(s []byte, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	return b.packet.WriteBytesWithShortSize(s, a)
}

func (b *Buffer) WriteBytesWithSize(s []byte, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	return b.packet.WriteBytesWithSize(s, a)
}

func (b *Buffer) ToByteArray() []byte {
	buffer := bytes.CreateBuilder()
	buffer.WriteShort(b.tlvVer)
	body := b.packet.Bytes()
	buffer.WriteShort(len(body))
	buffer.WriteBytes(body)
	return buffer.Bytes()
}
