package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"util/hex"
)

type Packet struct {
	buffer *bytes.Buffer
}

func (packet *Packet) WriteByte(i int) error {
	j := int8(i)
	return binary.Write(packet.buffer, binary.BigEndian, &j)
}

func (packet *Packet) WriteString(s string) error {
	_, err := packet.buffer.WriteString(s)
	return err
}

func (packet *Packet) WriteBytesWithSize(data []byte, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	packet.WriteInt(len(data) + a)
	return packet.WriteBytes(data)
}

func (packet *Packet) WriteBytesWithShortSize(data []byte, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	packet.WriteShort(len(data) + a)
	return packet.WriteBytes(data)
}

func (packet *Packet) WriteStringWithShortSize(s string, add ...int) error {
	data := []byte(s)
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	return packet.WriteBytesWithShortSize(data, a)
}

func (packet *Packet) WriteStringWithSize(s string, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	return packet.WriteBytesWithSize([]byte(s), a)
}

func (packet *Packet) WriteBytes(bs []byte) error {
	_, err := packet.buffer.Write(bs)
	return err
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

// CreateBuilderByData 通过字节组创建字节构建器
func CreateBuilderByData(data []byte) *Packet {
	return &Packet{buffer: bytes.NewBuffer(data)}
}

// CreateBuilder 创建字节构建器
func CreateBuilder() *Packet {
	return &Packet{buffer: bytes.NewBuffer([]byte{})}
}

func ToByteArray(str string) []byte {
	return []byte(str)
}

func BufToInt32(b []byte) (int, error) {
	if b != nil {
		if len(b) == 3 {
			b = append([]byte{0}, b...)
		}
		bytesBuffer := bytes.NewBuffer(b)
		switch len(b) {
		case 1:
			var tmp int8
			err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
			return int(tmp), err
		case 2:
			var tmp int16
			err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
			return int(tmp), err
		case 4:
			var tmp int32
			err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
			return int(tmp), err
		default:
			return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
		}
	}
	return 0, errors.New("Can't convert")
}

func Int32ToBuf(i int) []byte {
	tmp := int32(i)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &tmp)
	return bytesBuffer.Bytes()
}
