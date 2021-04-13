package jce

import (
	"bytes"
	"encoding/binary"
	"util/jce/codec"
	"util/jce/requestf"
)

type Jce struct {
	requestId            int
	data                 map[string][]byte
	servantName, funName string
}

func NewJce() *Jce {
	return &Jce{
		data: map[string][]byte{},
	}
}

func ParseJce(data []byte) *Jce {
	jce := &Jce{
		data: map[string][]byte{},
	}
	response := &requestf.RequestPacket{}
	err := response.ReadFrom(codec.NewReader(data))
	if err == nil {
		buffer := make([]byte, len(response.SBuffer))
		for i, datum := range response.SBuffer {
			buffer[i] = byte(datum)
		}
		reader := codec.NewReader(buffer)
		m, _ := reader.ReadStringBytesMap(0)
		jce.data = m
		jce.requestId = int(response.IRequestId)
		jce.funName = response.SFuncName
		jce.servantName = response.SServantName
	}
	return jce
}

func (j *Jce) PutData(key string, data []byte) {
	j.data[key] = data
}

func (j *Jce) GetData(key string) *codec.Reader {
	data, ok := j.data[key]
	if ok {
		return codec.NewReader(data)
	}
	return nil
}

func (j *Jce) SetRequestId(requestId int) {
	j.requestId = requestId
}

func (j *Jce) GetRequestId() int {
	return j.requestId
}

func (j *Jce) Bytes() []byte {
	req := requestf.RequestPacket{}
	req.IVersion = 3
	req.CPacketType = 0
	req.IMessageType = 0
	req.IRequestId = int32(j.requestId)
	req.SServantName = j.servantName
	req.SFuncName = j.funName

	park := codec.NewBuffer()
	park.WriteStringBytesMap(j.data, 0)

	data := park.ToBytes()
	buffer := make([]int8, len(data))
	for i, datum := range data {
		buffer[i] = int8(datum)
	}

	req.SBuffer = buffer
	req.ITimeout = 0
	req.Context = map[string]string{}
	req.Status = map[string]string{}
	pack, err := func() ([]byte, error) {
		sbuf := bytes.NewBuffer(nil)
		sbuf.Write(make([]byte, 4))
		os := codec.NewBuffer()
		err := req.WriteTo(os)
		if err != nil {
			return nil, err
		}
		bs := os.ToBytes()
		sbuf.Write(bs)
		l := sbuf.Len()
		binary.BigEndian.PutUint32(sbuf.Bytes(), uint32(l))
		return sbuf.Bytes(), nil
	}()
	if err == nil {
		return pack
	}
	return nil
}
