package bluesky

import (
	"container/list"
	"util/hex"
	"util/tcp"
)

type Receiver struct {
	client *tcp.Tcp

	waiter *list.List
}

type Packet struct {
	Cmd string
	Seq int

	Uin  int
	Body []byte

	isOver bool

	Element *list.Element
}

// InitReceive 初始化借包器
func (r *BlueSky) InitReceive() *Receiver {
	receiver := Receiver{
		client: r.client,
		waiter: list.New(),
	}
	r.client.Receive(func() func(body []byte) {
		return func(body []byte) {
			println(hex.Bytes2Str(body))

			l := receiver.waiter
			for element := l.Front(); element != nil; element = element.Next() {
				waiter, ok := element.Value.(*Packet)
				if ok {
					println(waiter.Cmd)

				}
			}
		}
	}())
	return &receiver
}
