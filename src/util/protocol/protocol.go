package protocol

import (
	"bytes"
	"encoding/binary"
	"util/protocol/codec"
	"util/protocol/requestf"
)

var maxPackageLength int = 10485760

// SetMaxPackageLength sets the max length of tars packet
func SetMaxPackageLength(len int) {
	maxPackageLength = len
}

func TarsRequest(rev []byte) (int, int) {
	if len(rev) < 4 {
		return 0, PackageLess
	}
	iHeaderLen := int(binary.BigEndian.Uint32(rev[0:4]))
	if iHeaderLen < 4 || iHeaderLen > maxPackageLength {
		return 0, PackageError
	}
	if len(rev) < iHeaderLen {
		return 0, PackageLess
	}
	return iHeaderLen, PackageFull
}

type TarsProtocol struct{}

func (p *TarsProtocol) RequestPack(req *requestf.RequestPacket) ([]byte, error) {
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

}
func (p *TarsProtocol) ResponseUnpack(pkg []byte) (*requestf.ResponsePacket, error) {
	packet := &requestf.ResponsePacket{}
	err := packet.ReadFrom(codec.NewReader(pkg[4:]))
	return packet, err
}
func (p *TarsProtocol) ParsePackage(rev []byte) (int, int) {
	return TarsRequest(rev)
}
