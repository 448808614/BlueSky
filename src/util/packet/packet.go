package packet

import (
	"bytes"
	"encoding/binary"
	"util/hex"
)

type Packet struct {
	buffer *bytes.Buffer
}

func (packet *Packet) WriteByte(i int) error {
	j := int8(i)
	return binary.Write(packet.buffer, binary.BigEndian, &j)
}

func (packet *Packet) WriteBoolean(z bool) error {
	var tmp byte = 0
	if z {
		tmp = 1
	}
	return packet.WriteByte(int(tmp))
}

func (packet *Packet) WriteShort(i int) error {
	j := int16(i)
	return binary.Write(packet.buffer, binary.BigEndian, &j)
}

func (packet *Packet) WriteInt(i int) error {
	j := int32(i)
	return binary.Write(packet.buffer, binary.BigEndian, &j)
}

func (packet *Packet) WriteLong(i int) error {
	j := int64(i)
	return binary.Write(packet.buffer, binary.BigEndian, &j)
}

func (packet *Packet) Cap() int {
	return packet.buffer.Cap()
}

func (packet *Packet) Size() int {
	return packet.buffer.Len()
}

func (packet *Packet) Bytes() []byte {
	return packet.buffer.Bytes()
}

func (packet *Packet) Hex() string {
	return hex.Bytes2Str(packet.Bytes())
}

// 通过字节组创建字节构建器
func CreateBuilderByData(data []byte) *Packet {
	return &Packet{buffer: bytes.NewBuffer(data)}
}

// 创建字节构建器
func CreateBuilder() *Packet {
	return &Packet{buffer: bytes.NewBuffer([]byte{})}
}
