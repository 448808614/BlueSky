package bluesky

import (
	"container/list"
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

func (r *Receiver) AddWaiter(cmd string, seq int) *list.Element {
	pack := Packet{Cmd: cmd, Seq: seq, isOver: false}
	return r.waiter.PushFront(&pack)
}

func (r *Receiver) WaitPacket(elem *list.Element) *Packet {
	waiter, ok := elem.Value.(*Packet)
	if ok {
		for waiter.isOver == false {
			if waiter.isOver == true {
				break
			}
		}
		return waiter
	} else {
		panic("这个屑玩意不是一个接包器~")
	}
}

// InitReceive 初始化借包器
func (s *BlueSky) InitReceive() *Receiver {
	receiver := Receiver{
		client: s.client,
		waiter: list.New(),
	}
	s.client.Receive(func() func(body []byte) {
		return func(body []byte) {

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
