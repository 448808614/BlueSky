package requestf

import (
	"fmt"
	"util/protocol/codec"
)

//ResponsePacket strcut implement
type ResponsePacket struct {
	IVersion     int16             `json:"iVersion"`
	CPacketType  int8              `json:"cPacketType"`
	IRequestId   int32             `json:"iRequestId"`
	IMessageType int32             `json:"iMessageType"`
	IRet         int32             `json:"iRet"`
	SBuffer      []int8            `json:"sBuffer"`
	Status       map[string]string `json:"status"`
	SResultDesc  string            `json:"sResultDesc"`
	Context      map[string]string `json:"context"`
}

func (st *ResponsePacket) ResetDefault() {
	st.CPacketType = 0
	st.IMessageType = 0
	st.IRet = 0
}

//ReadFrom reads  from _is and put into struct.
func (st *ResponsePacket) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.ReadInt16(&st.IVersion, 1, true)
	if err != nil {
		return err
	}

	err = _is.ReadInt8(&st.CPacketType, 2, true)
	if err != nil {
		return err
	}

	err = _is.ReadInt32(&st.IRequestId, 3, true)
	if err != nil {
		return err
	}

	err = _is.ReadInt32(&st.IMessageType, 4, true)
	if err != nil {
		return err
	}

	err = _is.ReadInt32(&st.IRet, 5, true)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(6, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}
		st.SBuffer = make([]int8, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = _is.ReadInt8(&st.SBuffer[i0], 0, false)
			if err != nil {
				return err
			}
		}
	} else if ty == codec.SimpleList {

		err, _ = _is.SkipTo(codec.BYTE, 0, true)
		if err != nil {
			return err
		}
		err = _is.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}
		err = _is.ReadSliceInt8(&st.SBuffer, length, true)
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}
	}

	err, have = _is.SkipTo(codec.MAP, 7, true)
	if err != nil {
		return err
	}

	err = _is.ReadInt32(&length, 0, true)
	if err != nil {
		return err
	}
	st.Status = make(map[string]string)
	for i1, e1 := int32(0), length; i1 < e1; i1++ {
		var k1 string
		var v1 string

		err = _is.ReadString(&k1, 0, false)
		if err != nil {
			return err
		}

		err = _is.ReadString(&v1, 1, false)
		if err != nil {
			return err
		}

		st.Status[k1] = v1
	}

	err = _is.ReadString(&st.SResultDesc, 8, false)
	if err != nil {
		return err
	}

	err, have = _is.SkipTo(codec.MAP, 9, false)
	if err != nil {
		return err
	}
	if have {
		err = _is.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}
		st.Context = make(map[string]string)
		for i2, e2 := int32(0), length; i2 < e2; i2++ {
			var k2 string
			var v2 string

			err = _is.ReadString(&k2, 0, false)
			if err != nil {
				return err
			}

			err = _is.ReadString(&v2, 1, false)
			if err != nil {
				return err
			}

			st.Context[k2] = v2
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *ResponsePacket) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ResponsePacket, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *ResponsePacket) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.WriteInt16(st.IVersion, 1)
	if err != nil {
		return err
	}

	err = _os.WriteInt8(st.CPacketType, 2)
	if err != nil {
		return err
	}

	err = _os.WriteInt32(st.IRequestId, 3)
	if err != nil {
		return err
	}

	err = _os.WriteInt32(st.IMessageType, 4)
	if err != nil {
		return err
	}

	err = _os.WriteInt32(st.IRet, 5)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.SimpleList, 6)
	if err != nil {
		return err
	}
	err = _os.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}
	err = _os.WriteInt32(int32(len(st.SBuffer)), 0)
	if err != nil {
		return err
	}
	err = _os.WriteSliceInt8(st.SBuffer)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.MAP, 7)
	if err != nil {
		return err
	}
	err = _os.WriteInt32(int32(len(st.Status)), 0)
	if err != nil {
		return err
	}
	for k3, v3 := range st.Status {

		err = _os.WriteString(k3, 0)
		if err != nil {
			return err
		}

		err = _os.WriteString(v3, 1)
		if err != nil {
			return err
		}
	}

	err = _os.WriteString(st.SResultDesc, 8)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.MAP, 9)
	if err != nil {
		return err
	}
	err = _os.WriteInt32(int32(len(st.Context)), 0)
	if err != nil {
		return err
	}
	for k4, v4 := range st.Context {

		err = _os.WriteString(k4, 0)
		if err != nil {
			return err
		}

		err = _os.WriteString(v4, 1)
		if err != nil {
			return err
		}
	}

	return nil
}

//WriteBlock encode struct
func (st *ResponsePacket) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}
