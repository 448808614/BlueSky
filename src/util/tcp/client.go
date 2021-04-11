package tcp

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"util/packet"
)

type Tcp struct {
	conn      *net.TCPConn
	isReceive bool
	reader    *bufio.Reader
}

func CreateTcp(host string, port int) *Tcp {
	conn, err := createConn(host, port)
	if err == nil {
		tcp := Tcp{
			conn:      conn,
			isReceive: false,
			reader:    nil,
		}
		return &tcp
	}
	log.Default().Println("Tcp connect failed.", err)
	return nil
}

func (t *Tcp) Send(body []byte) {
	go func() {
		_, err := t.conn.Write(body)
		if err != nil {
			log.Default().Println(err)
		}
	}()
}

func (t *Tcp) Receive(block func(body []byte)) {
	if !t.isReceive {
		go func() {
			if t.reader == nil {
				t.reader = bufio.NewReader(t.conn)
			}
			for {
				var ll = make([]byte, 4)
				_, err := t.reader.Read(ll)
				if err == nil {
					l, err := packet.BufToInt32(ll)
					if err == nil {
						var bb = make([]byte, l)
						_, err := t.reader.Read(bb)
						if err == nil {
							go block(bb)
						} else {
							log.Default().Println(err)
						}
					} else {
						log.Default().Println(err)
					}
				} else {
					log.Default().Println(err)
				}
			}
		}()
	} else {
		log.Default().Println("Receiver has been injected")
	}
}

func createConn(host string, port int) (*net.TCPConn, error) {
	addr, err := net.ResolveTCPAddr("tcp", host+":"+strconv.Itoa(port))
	if err == nil {
		conn, err := net.DialTCP("tcp", nil, addr)
		if err == nil {
			return conn, nil
		}
		return nil, err
	}
	return nil, err
}
