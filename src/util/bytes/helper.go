package bytes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
)

func RandBytes(size int) []byte {
	key := make([]byte, size)
	_, _ = rand.Read(key)
	return key
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
