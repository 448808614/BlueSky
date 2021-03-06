// Package codec implement
// 支持tars2go的底层库，用于基础类型的序列化
// 高级类型的序列化，由代码生成器，转换为基础类型的序列化

package codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"unsafe"
)

//jce type
const (
	BYTE byte = iota
	SHORT
	INT
	LONG
	FLOAT
	DOUBLE
	STRING1
	STRING4
	MAP
	LIST
	StructBegin
	StructEnd
	ZeroTag
	SimpleList
)

var typeToStr = []string{
	"Byte",
	"Short",
	"Int",
	"Long",
	"Float",
	"Double",
	"String1",
	"String4",
	"Map",
	"List",
	"StructBegin",
	"StructEnd",
	"ZeroTag",
	"SimpleList",
}

func getTypeStr(t int) string {
	if t < len(typeToStr) {
		return typeToStr[t]
	}
	return "invalidType"
}

// Buffer is wrapper of bytes.Buffer
type Buffer struct {
	buf *bytes.Buffer
}

// Reader is wrapper of bytes.Reader
type Reader struct {
	ref []byte
	buf *bytes.Reader
}

//go:nosplit
func bWriteU8(w *bytes.Buffer, data uint8) error {
	err := w.WriteByte(data)
	return err
}

//go:nosplit
func bWriteU16(w *bytes.Buffer, data uint16) error {
	var b [2]byte
	var bs []byte
	bs = b[:]
	binary.BigEndian.PutUint16(bs, data)
	_, err := w.Write(bs)
	return err
}

//go:nosplit
func bWriteU32(w *bytes.Buffer, data uint32) error {
	var b [4]byte
	var bs []byte
	bs = b[:]
	binary.BigEndian.PutUint32(bs, data)
	_, err := w.Write(bs)
	return err
}

//go:nosplit
func bWriteU64(w *bytes.Buffer, data uint64) error {
	var b [8]byte
	var bs []byte
	bs = b[:]
	binary.BigEndian.PutUint64(bs, data)
	_, err := w.Write(bs)
	return err
}

//go:nosplit
func bReadU8(r *bytes.Reader, data *uint8) error {
	var err error
	*data, err = r.ReadByte()
	return err
}

//go:nosplit
func bReadU16(r *bytes.Reader, data *uint16) error {
	var b [2]byte
	var bs []byte
	bs = b[:]
	_, err := r.Read(bs)
	*data = binary.BigEndian.Uint16(bs)
	return err
}

//go:nosplit
func bReadU32(r *bytes.Reader, data *uint32) error {
	var b [4]byte
	var bs []byte
	bs = b[:]
	_, err := r.Read(bs)
	*data = binary.BigEndian.Uint32(bs)
	return err
}

//go:nosplit
func bReadU64(r *bytes.Reader, data *uint64) error {
	var b [8]byte
	var bs []byte
	bs = b[:]
	_, err := r.Read(bs)
	*data = binary.BigEndian.Uint64(bs)
	return err
}

func (os *Buffer) WriteHead(ty byte, tag byte) error {
	if tag < 15 {
		data := (tag << 4) | ty
		return os.buf.WriteByte(data)
	} else {
		data := (15 << 4) | ty
		if err := os.buf.WriteByte(data); err != nil {
			return err
		}
		return os.buf.WriteByte(tag)
	}
}

// Reset clean the buffer.
func (os *Buffer) Reset() {
	os.buf.Reset()
}

// WriteSliceUint8 write []uint8 to the buffer.
func (os *Buffer) WriteSliceUint8(data []uint8) error {
	_, err := os.buf.Write(data)
	return err
}

// WriteSliceInt8 write []int8 to the buffer.
func (os *Buffer) WriteSliceInt8(data []int8) error {
	_, err := os.buf.Write(*(*[]uint8)(unsafe.Pointer(&data)))
	return err
}

func (os *Buffer) WriteStringBytesMap(data map[string][]byte, tag byte) error {
	err := os.WriteHead(MAP, tag)
	if err != nil {
		return err
	}
	err = os.WriteInt32(int32(len(data)), 0)
	if err != nil {
		return err
	}
	for k, v := range data {
		err = os.WriteString(k, 0)
		if err != nil {
			return err
		}

		err = os.WriteHead(SimpleList, 1)
		if err != nil {
			return err
		}
		err = os.WriteHead(BYTE, 0)
		if err != nil {
			return err
		}
		err = os.WriteInt32(int32(len(v)), 0)
		if err != nil {
			return err
		}
		err = os.WriteBytes(v)
		if err != nil {
			return err
		}
	}

	return err
}

/*
func (os *Buffer) WriteMap(data map[interface{}]interface{}, tag byte) error {
	err := os.WriteHead(MAP, 0)
	if err != nil {
		return err
	}
	err = os.WriteInt32(int32(len(data)), 0)
	if err != nil {
		return err
	}
	for k, v := range data {
		switch vv := k.(type) {
		case string:
			err = os.WriteString(vv, 0)
			if err != nil {
				return err
			}
		case int64:
			err = os.WriteInt64(vv, 0)
			if err != nil {
				return err
			}
		case int32:
			err = os.WriteInt32(vv, 0)
			if err != nil {
				return err
			}
		case int16:
			err = os.WriteInt16(vv, 0)
			if err != nil {
				return err
			}
		case int8:
			err = os.WriteInt8(vv, 0)
			if err != nil {
				return err
			}
		case uint64:
			err = os.WriteUint32(uint32(vv), 0)
			if err != nil {
				return err
			}
		case uint32:
			err = os.WriteUint32(vv, 0)
			if err != nil {
				return err
			}
		case uint16:
			err = os.WriteUint16(vv, 0)
			if err != nil {
				return err
			}
		case uint8:
			err = os.WriteUint8(vv, 0)
			if err != nil {
				return err
			}
		}
		switch v.(type) {
		case

		}



		err = os.WriteHead(codec.SIMPLE_LIST, 1)
		if err != nil {
			return err
		}
		err = os.WriteHead(codec.BYTE, 0)
		if err != nil {
			return err
		}
		err = os.Write_int32(int32(len(v)), 0)
		if err != nil {
			return err
		}
		err = os.Write_bytes(v)
		if err != nil {
			return err
		}
	}

	return  err


}
实现起来太麻烦了，不实现了，emmmmm(懒
*/

// WriteBytes write []byte to the buffer
func (os *Buffer) WriteBytes(data []byte) error {
	_, err := os.buf.Write(data)
	return err
}

// WriteInt8 write int8 with the tag.
func (os *Buffer) WriteInt8(data int8, tag byte) error {
	var err error
	if data == 0 {
		if err = os.WriteHead(ZeroTag, tag); err != nil {
			return err
		}
	} else {
		if err = os.WriteHead(BYTE, tag); err != nil {
			return err
		}

		if err = os.buf.WriteByte(byte(data)); err != nil {
			return err
		}
	}
	return nil
}

// WriteUint8 write uint8 with the tag
func (os *Buffer) WriteUint8(data uint8, tag byte) error {
	return os.WriteInt16(int16(data), tag)
}

// WriteBool write bool with the tag.
func (os *Buffer) WriteBool(data bool, tag byte) error {
	tmp := int8(0)
	if data {
		tmp = 1
	}
	return os.WriteInt8(tmp, tag)
}

// WriteInt16 writes the int16 with the tag.
func (os *Buffer) WriteInt16(data int16, tag byte) error {
	var err error
	if data >= math.MinInt8 && data <= math.MaxInt8 {
		if err = os.WriteInt8(int8(data), tag); err != nil {
			return err
		}
	} else {
		if err = os.WriteHead(SHORT, tag); err != nil {
			return err
		}

		if err = bWriteU16(os.buf, uint16(data)); err != nil {
			return err
		}
	}
	return nil
}

// WriteUint16 write uint16 with the tag.
func (os *Buffer) WriteUint16(data uint16, tag byte) error {
	return os.WriteInt32(int32(data), tag)
}

// WriteInt32 write int32 with the tag.
func (os *Buffer) WriteInt32(data int32, tag byte) error {
	var err error
	if data >= math.MinInt16 && data <= math.MaxInt16 {
		if err = os.WriteInt16(int16(data), tag); err != nil {
			return err
		}
	} else {
		if err = os.WriteHead(INT, tag); err != nil {
			return err
		}

		if err = bWriteU32(os.buf, uint32(data)); err != nil {
			return err
		}
	}
	return nil
}

// WriteUint32 write uint32 data with the tag.
func (os *Buffer) WriteUint32(data uint32, tag byte) error {
	return os.WriteInt64(int64(data), tag)
}

// WriteInt64 write int64 with the tag.
func (os *Buffer) WriteInt64(data int64, tag byte) error {
	var err error
	if data >= math.MinInt32 && data <= math.MaxInt32 {
		if err = os.WriteInt32(int32(data), tag); err != nil {
			return err
		}
	} else {
		if err = os.WriteHead(LONG, tag); err != nil {
			return err
		}

		if err = bWriteU64(os.buf, uint64(data)); err != nil {
			return err
		}
	}
	return nil
}

// WriteFloat32 writes float32 with the tag.
func (os *Buffer) WriteFloat32(data float32, tag byte) error {
	var err error
	if err = os.WriteHead(FLOAT, tag); err != nil {
		return err
	}

	err = bWriteU32(os.buf, math.Float32bits(data))
	return err
}

// WriteFloat64 writes float64 with the tag.
func (os *Buffer) WriteFloat64(data float64, tag byte) error {
	var err error
	if err = os.WriteHead(DOUBLE, tag); err != nil {
		return err
	}

	err = bWriteU64(os.buf, math.Float64bits(data))
	return err
}

// WriteString writes string data with the tag.
func (os *Buffer) WriteString(data string, tag byte) error {
	var err error
	if len(data) > 255 {
		if err = os.WriteHead(STRING4, tag); err != nil {
			return err
		}

		if err = bWriteU32(os.buf, uint32(len(data))); err != nil {
			return err
		}
	} else {
		if err = os.WriteHead(STRING1, tag); err != nil {
			return err
		}

		if err = bWriteU8(os.buf, byte(len(data))); err != nil {
			return err
		}
	}

	if _, err = os.buf.WriteString(data); err != nil {
		return err
	}
	return nil
}

// ToBytes make the buffer to []byte
func (os *Buffer) ToBytes() []byte {
	return os.buf.Bytes()
}

// Grow grows the size of the buffer.
func (os *Buffer) Grow(size int) {
	os.buf.Grow(size)
}

//Reset clean the Reader.
func (b *Reader) Reset(data []byte) {
	b.buf.Reset(data)
	b.ref = data
}

//go:nosplit
func (b *Reader) readHead() (ty, tag byte, err error) {
	data, err := b.buf.ReadByte()
	if err != nil {
		return
	}
	ty = data & 0x0f
	tag = (data & 0xf0) >> 4
	if tag == 15 {
		data, err = b.buf.ReadByte()
		if err != nil {
			return
		}
		tag = data
	}
	return
}

// unreadHead 回退一个head byte， curTag为当前读到的tag信息，当tag超过4位时则回退两个head byte
// unreadHead put back the current head byte.
func (b *Reader) unreadHead(curTag byte) {
	_ = b.buf.UnreadByte()
	if curTag >= 15 {
		_ = b.buf.UnreadByte()
	}
}

// Next return the []byte of next n .
//go:nosplit
func (b *Reader) Next(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	beg := len(b.ref) - b.buf.Len()
	_, _ = b.buf.Seek(int64(n), io.SeekCurrent)
	end := len(b.ref) - b.buf.Len()
	return b.ref[beg:end]
}

// Skip Skip the next n byte.
//go:nosplit
func (b *Reader) Skip(n int) {
	if n <= 0 {
		return
	}
	_, _ = b.buf.Seek(int64(n), io.SeekCurrent)
}

func (b *Reader) skipFieldMap() error {
	var l int32
	err := b.ReadInt32(&l, 0, true)
	if err != nil {
		return err
	}

	for i := int32(0); i < l*2; i++ {
		tyCur, _, err := b.readHead()
		if err != nil {
			return err
		}
		_ = b.skipField(tyCur)
	}
	return nil
}
func (b *Reader) skipFieldList() error {
	var l int32
	err := b.ReadInt32(&l, 0, true)
	if err != nil {
		return err
	}
	for i := int32(0); i < l; i++ {
		tyCur, _, err := b.readHead()
		if err != nil {
			return err
		}
		_ = b.skipField(tyCur)
	}
	return nil
}
func (b *Reader) skipFieldSimpleList() error {
	tyCur, _, err := b.readHead()
	if tyCur != BYTE {
		return fmt.Errorf("simple list need byte head. but get %d", tyCur)
	}
	if err != nil {
		return err
	}
	var l int32
	err = b.ReadInt32(&l, 0, true)
	if err != nil {
		return err
	}

	b.Skip(int(l))
	return nil
}

func (b *Reader) skipField(ty byte) error {
	switch ty {
	case BYTE:
		b.Skip(1)
		break
	case SHORT:
		b.Skip(2)
		break
	case INT:
		b.Skip(4)
		break
	case LONG:
		b.Skip(8)
		break
	case FLOAT:
		b.Skip(4)
		break
	case DOUBLE:
		b.Skip(8)
		break
	case STRING1:
		data, err := b.buf.ReadByte()
		if err != nil {
			return err
		}
		l := int(data)
		b.Skip(l)
		break
	case STRING4:
		var l uint32
		err := bReadU32(b.buf, &l)
		if err != nil {
			return err
		}
		b.Skip(int(l))
		break
	case MAP:
		err := b.skipFieldMap()
		if err != nil {
			return err
		}
		break
	case LIST:
		err := b.skipFieldList()
		if err != nil {
			return err
		}
		break
	case SimpleList:
		err := b.skipFieldSimpleList()
		if err != nil {
			return err
		}
		break
	case StructBegin:
		err := b.SkipToStructEnd()
		if err != nil {
			return err
		}
		break
	case StructEnd:
		break
	case ZeroTag:
		break
	default:
		return fmt.Errorf("invalid type")
	}
	return nil
}

// SkipToStructEnd for skip to the STRUCT_END tag.
func (b *Reader) SkipToStructEnd() error {
	for {
		ty, _, err := b.readHead()
		if err != nil {
			return err
		}

		err = b.skipField(ty)
		if err != nil {
			return err
		}
		if ty == StructEnd {
			break
		}
	}
	return nil
}

// SkipToNoCheck for skip to the none STRUCT_END tag.
func (b *Reader) SkipToNoCheck(tag byte, require bool) (error, bool, byte) {
	for {
		tyCur, tagCur, err := b.readHead()
		if err != nil {
			if require {
				return fmt.Errorf("Can not find Tag %d. But require. %s", tag, err.Error()), false, tyCur
			}
			return nil, false, tyCur
		}
		if tyCur == StructEnd || tagCur > tag {
			if require {
				return fmt.Errorf("Can not find Tag %d. But require. tagCur: %d, tyCur: %d",
					tag, tagCur, tyCur), false, tyCur
			}
			// 多读了一个head, 退回去
			b.unreadHead(tagCur)
			return nil, false, tyCur
		}
		if tagCur == tag {
			return nil, true, tyCur
		}
		// tagCur < tag
		if err = b.skipField(tyCur); err != nil {
			return err, false, tyCur
		}
	}
}

// SkipTo skip to the given tag.
func (b *Reader) SkipTo(ty, tag byte, require bool) (error, bool) {
	err, have, tyCur := b.SkipToNoCheck(tag, require)
	if err != nil {
		return err, false
	}
	if have && ty != tyCur {
		return fmt.Errorf("type not match, need %d, bug %d", ty, tyCur), false
	}
	return nil, have
}

// ReadSliceInt8 reads []int8 for the given length and the require or optional sign.
func (b *Reader) ReadSliceInt8(data *[]int8, len int32, require bool) error {
	if len <= 0 {
		return nil
	}

	*data = make([]int8, len)
	_, err := b.buf.Read(*(*[]uint8)(unsafe.Pointer(data)))
	if err != nil {
		err = fmt.Errorf("Read_slice_int8 error:%v", err)
	}
	return err
}

// ReadSliceUint8 reads []uint8 fore the given length and the require or optional sign.
func (b *Reader) ReadSliceUint8(data *[]uint8, len int32, require bool) error {
	if len <= 0 {
		return nil
	}

	*data = make([]uint8, len)
	_, err := b.buf.Read(*data)
	if err != nil {
		err = fmt.Errorf("Read_slice_uint8 error:%v", err)
	}
	return err
}

func (b *Reader) ReadStringBytesMap(tag byte) (map[string][]byte, error) {
	ret := map[string][]byte{}
	err, have := b.SkipTo(MAP, tag, false)
	if err != nil {
		return ret, err
	}
	var length int32 = 0
	err = b.ReadInt32(&length, 0, true)
	if err != nil {
		return ret, err
	}
	var ty byte
	for i, e := int32(0), length; i < e; i++ {
		var k string
		var v []byte
		err = b.ReadString(&k, 0, false)
		if err != nil {
			return ret, err
		}
		err, have, ty = b.SkipToNoCheck(1, false)
		if err != nil {
			return ret, err
		}
		if have {
			if ty == SimpleList {
				err, _ = b.SkipTo(BYTE, 0, true)
				if err != nil {
					return ret, err
				}
				var byteLen int32 = 0
				err = b.ReadInt32(&byteLen, 0, true)
				if err != nil {
					return ret, err
				}
				err = b.ReadBytes(&v, byteLen, true)
				if err != nil {
					return ret, err
				}
				ret[k] = v
			} else {
				err = fmt.Errorf("require vector, but not")
				return ret, err
			}
		}
	}
	return ret, err
}

//ReadBytes reads []byte for the given length and the require or optional sign.
func (b *Reader) ReadBytes(data *[]byte, len int32, require bool) error {
	*data = make([]byte, len)
	_, err := b.buf.Read(*data)
	return err
}

// ReadInt8 reads the int8 data for the tag and the require or optional sign.
func (b *Reader) ReadInt8(data *int8, tag byte, require bool) error {
	err, have, ty := b.SkipToNoCheck(tag, require)
	if err != nil {
		return err
	}
	if !have {
		return nil
	}
	switch ty {
	case ZeroTag:
		*data = 0
	case BYTE:
		var tmp uint8
		err = bReadU8(b.buf, &tmp)
		*data = int8(tmp)
	default:
		return fmt.Errorf("read 'int8' type mismatch, tag:%d, get type:%s", tag, getTypeStr(int(ty)))
	}
	if err != nil {
		err = fmt.Errorf("ReadInt8 tag:%d error:%v", tag, err)
	}
	return err
}

// ReadUint8 reads the uint8 for the tag and the require or optional sign.
func (b *Reader) ReadUint8(data *uint8, tag byte, require bool) error {
	n := int16(*data)
	err := b.ReadInt16(&n, tag, require)
	*data = uint8(n)
	return err
}

// ReadBool reads the bool value for the tag and the require or optional sign.
func (b *Reader) ReadBool(data *bool, tag byte, require bool) error {
	var tmp int8
	err := b.ReadInt8(&tmp, tag, require)
	if err != nil {
		return err
	}
	if tmp == 0 {
		*data = false
	} else {
		*data = true
	}
	return nil
}

// ReadInt16 reads the int16 value for the tag and the require or optional sign.
func (b *Reader) ReadInt16(data *int16, tag byte, require bool) error {
	err, have, ty := b.SkipToNoCheck(tag, require)
	if err != nil {
		return err
	}
	if !have {
		return nil
	}
	switch ty {
	case ZeroTag:
		*data = 0
	case BYTE:
		var tmp uint8
		err = bReadU8(b.buf, &tmp)
		*data = int16(int8(tmp))
	case SHORT:
		var tmp uint16
		err = bReadU16(b.buf, &tmp)
		*data = int16(tmp)
	default:
		return fmt.Errorf("read 'int16' type mismatch, tag:%d, get type:%s", tag, getTypeStr(int(ty)))
	}
	if err != nil {
		err = fmt.Errorf("Read_int16 tag:%d error:%v", tag, err)
	}
	return err
}

// ReadUint16 reads the uint16 value for the tag and the require or optional sign.
func (b *Reader) ReadUint16(data *uint16, tag byte, require bool) error {
	n := int32(*data)
	err := b.ReadInt32(&n, tag, require)
	*data = uint16(n)
	return err
}

// ReadInt32 reads the int32 value for the tag and the require or optional sign.
func (b *Reader) ReadInt32(data *int32, tag byte, require bool) error {
	err, have, ty := b.SkipToNoCheck(tag, require)
	if err != nil {
		return err
	}
	if !have {
		return nil
	}
	switch ty {
	case ZeroTag:
		*data = 0
	case BYTE:
		var tmp uint8
		err = bReadU8(b.buf, &tmp)
		*data = int32(int8(tmp))
	case SHORT:
		var tmp uint16
		err = bReadU16(b.buf, &tmp)
		*data = int32(int16(tmp))
	case INT:
		var tmp uint32
		err = bReadU32(b.buf, &tmp)
		*data = int32(tmp)
	default:
		return fmt.Errorf("read 'int32' type mismatch, tag:%d, get type:%s", tag, getTypeStr(int(ty)))
	}
	if err != nil {
		err = fmt.Errorf("ReadInt32 tag:%d error:%v", tag, err)
	}
	return err
}

// ReadUint32 reads the uint32 value for the tag and the require or optional sign.
func (b *Reader) ReadUint32(data *uint32, tag byte, require bool) error {
	n := int64(*data)
	err := b.ReadInt64(&n, tag, require)
	*data = uint32(n)
	return err
}

// ReadInt64 reads the int64 value for the tag and the require or optional sign.
func (b *Reader) ReadInt64(data *int64, tag byte, require bool) error {
	err, have, ty := b.SkipToNoCheck(tag, require)
	if err != nil {
		return err
	}
	if !have {
		return nil
	}
	switch ty {
	case ZeroTag:
		*data = 0
	case BYTE:
		var tmp uint8
		err = bReadU8(b.buf, &tmp)
		*data = int64(int8(tmp))
	case SHORT:
		var tmp uint16
		err = bReadU16(b.buf, &tmp)
		*data = int64(int16(tmp))
	case INT:
		var tmp uint32
		err = bReadU32(b.buf, &tmp)
		*data = int64(int32(tmp))
	case LONG:
		var tmp uint64
		err = bReadU64(b.buf, &tmp)
		*data = int64(tmp)
	default:
		return fmt.Errorf("read 'int64' type mismatch, tag:%d, get type:%s", tag, getTypeStr(int(ty)))
	}
	if err != nil {
		err = fmt.Errorf("Read_int64 tag:%d error:%v", tag, err)
	}

	return err
}

// ReadFloat32 reads the float32 value for the tag and the require or optional sign.
func (b *Reader) ReadFloat32(data *float32, tag byte, require bool) error {
	err, have, ty := b.SkipToNoCheck(tag, require)
	if err != nil {
		return err
	}
	if !have {
		return nil
	}

	switch ty {
	case ZeroTag:
		*data = 0
	case FLOAT:
		var tmp uint32
		err = bReadU32(b.buf, &tmp)
		*data = math.Float32frombits(tmp)
	default:
		return fmt.Errorf("read 'float' type mismatch, tag:%d, get type:%s", tag, getTypeStr(int(ty)))
	}

	if err != nil {
		err = fmt.Errorf("Read_float32 tag:%d error:%v", tag, err)
	}
	return err
}

// ReadFloat64 reads the float64 value for the tag and the require or optional sign.
func (b *Reader) ReadFloat64(data *float64, tag byte, require bool) error {
	err, have, ty := b.SkipToNoCheck(tag, require)
	if err != nil {
		return err
	}
	if !have {
		return nil
	}

	switch ty {
	case ZeroTag:
		*data = 0
	case FLOAT:
		var tmp uint32
		err = bReadU32(b.buf, &tmp)
		*data = float64(math.Float32frombits(tmp))
	case DOUBLE:
		var tmp uint64
		err = bReadU64(b.buf, &tmp)
		*data = math.Float64frombits(tmp)
	default:
		return fmt.Errorf("read 'double' type mismatch, tag:%d, get type:%s", tag, getTypeStr(int(ty)))
	}

	if err != nil {
		err = fmt.Errorf("Read_float64 tag:%d error:%v", tag, err)
	}
	return err
}

// ReadString reads the string value for the tag and the require or optional sign.
func (b *Reader) ReadString(data *string, tag byte, require bool) error {
	err, have, ty := b.SkipToNoCheck(tag, require)
	if err != nil {
		return err
	}
	if !have {
		return nil
	}

	if ty == STRING4 {
		var l uint32
		err = bReadU32(b.buf, &l)
		if err != nil {
			return fmt.Errorf("ReadString4 tag:%d error:%v", tag, err)
		}
		buff := b.Next(int(l))
		*data = string(buff)
	} else if ty == STRING1 {
		var l uint8
		err = bReadU8(b.buf, &l)
		if err != nil {
			return fmt.Errorf("ReadString1 tag:%d error:%v", tag, err)
		}
		buff := b.Next(int(l))
		*data = string(buff)
	} else {
		return fmt.Errorf("need string, tag:%d, but type is %s", tag, getTypeStr(int(ty)))
	}
	return nil
}

//ToString make the reader to string
func (b *Reader) ToString() string {
	return string(b.ref[:])
}

// ToBytes make the reader to string
func (b *Reader) ToBytes() []byte {
	return b.ref
}

// NewReader returns *Reader
func NewReader(data []byte) *Reader {
	return &Reader{buf: bytes.NewReader(data), ref: data}
}

// NewBuffer returns *Buffer
func NewBuffer() *Buffer {
	return &Buffer{buf: &bytes.Buffer{}}
}

// FromInt8 NewReader(FromInt8(vec))
func FromInt8(vec []int8) []byte {
	return *(*[]byte)(unsafe.Pointer(&vec))
}
