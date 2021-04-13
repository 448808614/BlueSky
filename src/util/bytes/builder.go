package bytes

import (
	"bytes"
	"encoding/binary"
	"util/hex"
)

type Builder struct {
	buffer *bytes.Buffer
}

func (b *Builder) WriteByte(i int) error {
	j := int8(i)
	return binary.Write(b.buffer, binary.BigEndian, &j)
}

func (b *Builder) WriteString(s string) error {
	_, err := b.buffer.WriteString(s)
	return err
}

func (b *Builder) WriteBytesWithSize(data []byte, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	b.WriteInt(len(data) + a)
	return b.WriteBytes(data)
}

func (b *Builder) WriteBytesWithShortSize(data []byte, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	b.WriteShort(len(data) + a)
	return b.WriteBytes(data)
}

func (b *Builder) WriteStringWithShortSize(s string, add ...int) error {
	data := []byte(s)
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	return b.WriteBytesWithShortSize(data, a)
}

func (b *Builder) WriteStringWithSize(s string, add ...int) error {
	a := 0
	if len(add) >= 1 {
		a = add[0]
	}
	return b.WriteBytesWithSize([]byte(s), a)
}

func (b *Builder) WriteBytes(bs []byte) error {
	_, err := b.buffer.Write(bs)
	return err
}

func (b *Builder) WriteBoolean(z bool) error {
	var tmp byte = 0
	if z {
		tmp = 1
	}
	return b.WriteByte(int(tmp))
}

func (b *Builder) WriteShort(i int) error {
	j := int16(i)
	return binary.Write(b.buffer, binary.BigEndian, &j)
}

func (b *Builder) WriteInt(i int) error {
	j := int32(i)
	return binary.Write(b.buffer, binary.BigEndian, &j)
}

func (b *Builder) WriteLong(i int) error {
	j := int64(i)
	return binary.Write(b.buffer, binary.BigEndian, &j)
}

func (b *Builder) Cap() int {
	return b.buffer.Cap()
}

func (b *Builder) Size() int {
	return b.buffer.Len()
}

func (b *Builder) Bytes() []byte {
	return b.buffer.Bytes()
}

func (b *Builder) Hex() string {
	return hex.Bytes2Str(b.Bytes())
}

// CreateBuilderByData 通过字节组创建字节构建器
func CreateBuilderByData(data []byte) *Builder {
	return &Builder{buffer: bytes.NewBuffer(data)}
}

// CreateBuilder 创建字节构建器
func CreateBuilder() *Builder {
	return &Builder{buffer: bytes.NewBuffer([]byte{})}
}
