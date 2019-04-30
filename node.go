package mvds

import (
	"github.com/status-im/mvds/securetransport"
	"github.com/status-im/mvds/storage"
)

type calculateSendTime func(count uint64, lastTime uint64) uint64
type PeerId [32]byte

type State struct {
	HoldFlag    bool
	AckFlag     bool
	RequestFlag bool
	SendCount   uint64
	SendTime    uint64
}

type Node struct {
	ms storage.MessageStore
	st securetransport.Node

	ss map[MessageID]State

	queue map[PeerId]Payload // @todo we use this so we can queue messages rather than sending stuff alone
							 // @todo make this a new object which is mutexed

	sc calculateSendTime

	id []byte // @todo

	Send     <-chan []byte
	Received chan<- []byte // @todo message type
}

func (n *Node) Start() error {

	// @todo start listening to both the send channel and what the transport receives for later handling

	return nil
}

func (n *Node) onPayload(payload Payload) {
	// @todo probably needs to check that its not null and all that
	// @todo do these need to be go routines?
	n.onAck(payload.ack)
	n.onRequest(payload.request)
	n.onOffer(payload.offer)

	for _, m := range payload.messages {
		n.onMessage(m)
	}
}

func (n *Node) onOffer(msg Offer) {

}

func (n *Node) onRequest(msg Request) {
	for _, id := range msg.Messages {
		_, err := n.ms.GetMessage(id)
		if err != nil {
			// @todo
		}

		n.send(id)
	}
}

func (n *Node) onAck(msg Ack) {
	for _, id := range msg.Messages {
		s, _ := n.ss[id]
		s.HoldFlag = true
	}
}

func (n *Node) onMessage(msg Message) {

	// @todo handle

	n.ss[msg.ID()] = State{
		HoldFlag: true,
		AckFlag: true,
		RequestFlag: false,
		SendTime: 0,
		SendCount: 0,
	}

	err := n.ms.SaveMessage(msg)
	if err != nil {
		// @todo process, should this function ever even have an error?
	}
}

func (n *Node) send(id MessageID) error {
	s, _ := n.ss[id]

	s.SendCount += 1
	s.SendTime = n.sc(s.SendCount, s.SendTime)

	// @todo actually send

	return nil
}


func (n *Node) sendForPeer(peer PeerId) {

}